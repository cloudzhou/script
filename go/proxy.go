package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

var proxyUrl string
var contextMap map[string]*Context = make(map[string]*Context)
var mutex sync.Mutex

func getContext(addrHex string) *Context {
	mutex.Lock()
	defer mutex.Unlock()
	return contextMap[addrHex]
}

func dumpContextMap() map[string][]string {
	mutex.Lock()
	defer mutex.Unlock()
	m := make(map[string][]string)
	for s, c := range contextMap {
		m[s] = c.addrs()
	}
	return m
}

func register(addrHex string, c *Context) {
	mutex.Lock()
	defer mutex.Unlock()
	contextMap[addrHex] = c
}

func unregister(c *Context) {
	mutex.Lock()
	defer mutex.Unlock()
	addrs := c.addrs()
	for _, addr := range addrs {
		if ct, ok := contextMap[addr]; ok && c == ct {
			delete(contextMap, addr)
		}
	}
}

// Header: X-MESH-META = Ver/Option/Flags/Proto, X-MESH-ADDR=DstAddr/SrcAddr, X-MESH-OPTIONLIST=hex(OptionList)
type Packet struct {
	Ver        int
	Option     int
	Flags      int
	Proto      int
	Len        int
	DstAddr    []byte
	SrcAddr    []byte
	OtLen      int
	OptionList []byte
	Data       []byte
}

func (p *Packet) unmarshalFromHttp(header http.Header, body io.ReadCloser) error {
	var err error
	var strs []string
	meshMeta := header.Get("X-MESH-META")
	strs = strings.Split(meshMeta, "/")
	if len(strs) != 4 {
		return errors.New("X-MESH-META must be Ver/Option/Flags/Proto")
	}
	var i int64
	i, _ = strconv.ParseInt(strs[0], 10, 0)
	p.Ver = int(i)
	i, _ = strconv.ParseInt(strs[1], 10, 0)
	p.Option = int(i)
	i, _ = strconv.ParseInt(strs[2], 10, 0)
	p.Flags = int(i)
	i, _ = strconv.ParseInt(strs[3], 10, 0)
	p.Proto = int(i)
	meshAddr := header.Get("X-MESH-ADDR")
	strs = strings.Split(meshAddr, "/")
	if len(strs) != 2 {
		return errors.New("X-MESH-ADDR must be DstAddr/SrcAddr")
	}
	p.DstAddr, err = hex.DecodeString(strs[0])
	if err != nil {
		return err
	}
	p.SrcAddr, err = hex.DecodeString(strs[1])
	if err != nil {
		return err
	}
	if p.Option == 1 {
		optionListHex := header.Get("X-MESH-OPTIONLIST")
		p.OptionList, err = hex.DecodeString(optionListHex)
		if err != nil {
			return err
		}
		p.OtLen = len(p.OptionList)
	}
	defer body.Close()
	p.Data, err = ioutil.ReadAll(body)
	p.Len = 16 + p.OtLen + len(p.Data)
	return err
}

func (p *Packet) marshal() ([]byte, error) {
	p.OtLen = len(p.OptionList)
	if p.OtLen > 0 {
		p.OtLen = p.OtLen + 2
	}
	p.Len = 16 + p.OtLen + len(p.Data)
	buf := make([]byte, p.Len)
	buf[0] = byte(p.Ver | p.Option<<2 | p.Flags<<3)
	buf[1] = byte(p.Proto)
	binary.LittleEndian.PutUint16(buf[2:4], uint16(p.Len))
	if len(p.DstAddr) != 6 || len(p.SrcAddr) != 6 {
		return nil, errors.New("len(p.DstAddr) and len(p.SrcAddr) must be 6")
	}
	copy(buf[4:10], p.DstAddr)
	copy(buf[10:16], p.SrcAddr)
	if p.OtLen > 0 {
		binary.LittleEndian.PutUint16(buf[16:18], uint16(p.OtLen))
		copy(buf[18:], p.OptionList)
	}
	copy(buf[16+p.OtLen:], p.Data)
	return buf, nil
}

func (p *Packet) dump() string {
	return fmt.Sprintf("Ver: %d, Option: %d, Flags: %d, Proto: %d, Len: %d, DstAddr: %s, SrcAddr: %s, OtLen: %d, OptionList: %s, Data: %s", p.Ver, p.Option, p.Flags, p.Proto, p.Len, hex.EncodeToString(p.DstAddr), hex.EncodeToString(p.SrcAddr), p.OtLen, hex.EncodeToString(p.OptionList), hex.EncodeToString(p.Data))
}

var packetPool = sync.Pool{
	New: func() interface{} {
		return &Packet{}
	},
}

func getPacket() *Packet {
	return packetPool.Get().(*Packet)
}

func putPacket(p *Packet) {
	p.DstAddr = nil
	p.SrcAddr = nil
	p.OptionList = nil
	p.Data = nil
	p.OtLen = 0
	p.Len = 0
	packetPool.Put(p)
}

type Context struct {
	conn    net.Conn
	reader  io.Reader
	addrMap map[string]interface{}
	sync.Mutex
}

func (c *Context) join(addrHex string) bool {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.addrMap[addrHex]; ok {
		return true
	}
	c.addrMap[addrHex] = struct{}{}
	return false
}

func (c *Context) leave(addrHex string) {
	c.Lock()
	defer c.Unlock()
	delete(c.addrMap, addrHex)
}

func (c *Context) addrs() []string {
	c.Lock()
	defer c.Unlock()
	addrs := make([]string, len(c.addrMap))
	i := 0
	for k := range c.addrMap {
		addrs[i] = k
		i++
	}
	return addrs
}

func (c *Context) read(packet *Packet) error {
	c.conn.SetReadDeadline(time.Now().Add(10 * time.Minute))
	buf := make([]byte, 4)
	_, err := io.ReadFull(c.reader, buf)
	if err != nil {
		return err
	}
	len := binary.LittleEndian.Uint16(buf[2:4])
	if len > 0xFFFF {
		return errors.New("Payload too much")
	}
	buf = append(buf, make([]byte, len-4)...)
	_, err = io.ReadFull(c.reader, buf[4:])
	if err != nil {
		return err
	}
	ord := int(buf[0])
	packet.Ver = (ord & 0x3)
	packet.Option = ((ord & 0x4) >> 2)
	packet.Flags = (ord >> 3)
	packet.Proto = int(buf[1])
	packet.Len = int(len)
	packet.DstAddr = buf[4:10]
	packet.SrcAddr = buf[10:16]
	if packet.Option == 0 {
		packet.Data = buf[16:]
		return nil
	}
	packet.OtLen = int(binary.LittleEndian.Uint16(buf[16:18]))
	if 16+packet.OtLen > packet.Len {
		return errors.New("16 + packet.OtLen > packet.Len")
	}
	packet.OptionList = buf[18 : 16+packet.OtLen]
	packet.Data = buf[16+packet.OtLen:]
	return nil
}

func (c *Context) write(packet *Packet) error {
	buf, err := packet.marshal()
	if err != nil {
		return err
	}
	c.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	_, err = c.conn.Write(buf)
	return err
}

type Handler interface {
	handle(conn net.Conn) error
}

type HttpHandler struct {
	client http.Client
}

func NewHttpHandler() *HttpHandler {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	return &HttpHandler{client: client}
}

func (h *HttpHandler) json(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(body)
}

func (h *HttpHandler) query(w http.ResponseWriter, c *Context) {
	addrs := c.addrs()
	m := map[string][]string{"nodes": addrs}
	body, _ := json.Marshal(m)
	h.json(w, body)
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	switch req.URL.Path {
	case "/deliver":
		p := getPacket()
		defer putPacket(p)
		err := p.unmarshalFromHttp(req.Header, req.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		dstAddrHex := hex.EncodeToString(p.DstAddr)
		c := getContext(dstAddrHex)
		if c == nil {
			http.Error(w, "dst addr node not registered", 404)
			return
		}
		err = c.write(p)
		if err != nil {
			http.Error(w, "context write error", 502)
			return
		}
		h.json(w, []byte(`{"status": 200, "message": "deliver success"}`))
	case "/mesh":
		node := values.Get("node")
		root := values.Get("root")
		action := values.Get("action")
		if action == "dump" {
			m := dumpContextMap()
			body, _ := json.Marshal(m)
			h.json(w, body)
			return
		}
		c := getContext(root)
		if c == nil {
			http.Error(w, "root node not registered", 404)
			return
		}
		switch action {
		case "query":
			h.query(w, c)
		case "join":
			c.join(node)
			h.query(w, c)
		case "leave":
			c.leave(node)
			h.query(w, c)
		default:
			http.Error(w, "action: dump/query/join/leave", 400)
			return
		}
	default:
		http.Error(w, "404 Not Found", 404)
		return
	}
}

func RunHttp() {
	httpHandler := NewHttpHandler()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        httpHandler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 5,
	}
	s.ListenAndServe()
}

func RunProxy() {
	l, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	proxyHandler := NewProxyHandler()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second)
			continue
		}
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println(r.(error))
				}
			}()
			err := proxyHandler.handle(conn)
			if err != nil && err != io.EOF {
				log.Println(err)
			}
		}()
	}
}

type ProxyHandler struct {
	client http.Client
}

func NewProxyHandler() *ProxyHandler {
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	return &ProxyHandler{client: client}
}

func (p *ProxyHandler) proxy(req *Packet, resp *Packet) error {
	httpReq, err := http.NewRequest("POST", proxyUrl, bytes.NewReader(req.Data))
	meshMeta := fmt.Sprintf("%d/%d/%d/%d", req.Ver, req.Option, req.Flags, req.Proto)
	httpReq.Header.Add("X-MESH-META", meshMeta)
	meshAddr := fmt.Sprintf("%s/%s", hex.EncodeToString(req.DstAddr), hex.EncodeToString(req.SrcAddr))
	httpReq.Header.Add("X-MESH-ADDR", meshAddr)
	if req.Option == 1 {
		optionListHex := hex.EncodeToString(req.OptionList)
		httpReq.Header.Add("X-MESH-OPTIONLIST", optionListHex)
	}
	httpReq.Header.Add("Content-Length", string(len(req.Data)))
	httpReq.Header.Add("Content-Type", "application/octet-stream")
	httpResp, err := p.client.Do(httpReq)
	if err != nil {
		return err
	}
	err = resp.unmarshalFromHttp(httpResp.Header, httpResp.Body)
	return err
}

func (p *ProxyHandler) handle(conn net.Conn) error {
	c := &Context{conn: conn, reader: bufio.NewReaderSize(conn, 4096), addrMap: make(map[string]interface{})}
	defer unregister(c)
	req := getPacket()
	defer putPacket(req)
	resp := getPacket()
	defer putPacket(resp)

	var err error
	for {
		err = c.read(req)
		if err != nil {
			return err
		}
		log.Println("req packet: {", req.dump(), "}")
		addrHex := hex.EncodeToString(req.SrcAddr)
		if !c.join(addrHex) {
			register(addrHex, c)
		}
		err = p.proxy(req, resp)
		if err != nil {
			return err
		}
		err = c.write(resp)
		if err != nil {
			return err
		}
	}
	return nil
}

func setProxyUrl() {
	flag.StringVar(&(proxyUrl), "proxyUrl", "", "proxy url")
	flag.Parse()
	if proxyUrl == "" {
		panic("need --proxyUrl parameter")
	}
}

func main() {
	setProxyUrl()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	go RunHttp()
	go RunProxy()
	<-sc
	log.Println("catch signal, exit")
}
