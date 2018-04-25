package main

import (
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type file struct {
	conn net.Conn
	name string
	c    chan struct{}
}

func (f *file) signalSend() {
	buf := make([]byte, 1)
	f.conn.Write(buf)
}

func (f *file) notify() {
	select {
	case f.c <- struct{}{}:
	default:
	}
}
func (f *file) wait() {
	select {
	case <-f.c:
	case <-time.After(time.Second * 60):
	}
}

var filemap map[uint32]*file = make(map[uint32]*file)
var mutex sync.Mutex

type Handler interface {
	handle(conn net.Conn) error
}

type HttpHandler struct {
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	randstr := req.URL.Path[1:]
	randint, err := strconv.ParseInt(randstr, 36, 64)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	mutex.Lock()
	file, ok := filemap[uint32(randint)]
	mutex.Unlock()
	if !ok {
		w.Write([]byte("file not found"))
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", file.name))
	file.signalSend()
	io.Copy(w, file.conn)
	file.notify()
}

func RunHttp() {
	httpHandler := NewHttpHandler()
	s := &http.Server{
		Addr:           ":7070",
		Handler:        httpHandler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 5,
	}
	s.ListenAndServe()
}

func RunProxy() {
	l, err := net.Listen("tcp", ":1024")
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
}

func NewProxyHandler() *ProxyHandler {
	return &ProxyHandler{}
}

func (p *ProxyHandler) handle(conn net.Conn) error {
	rint := uint32(rand.Int31())
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, rint)
	_, err := conn.Write(buf)
	if err != nil {
		panic(err)
	}
	log.Println(buf)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}
	l := int(binary.BigEndian.Uint32(buf))
	log.Println(l)
	if l > 1024 {
		return errors.New("name too large")
	}
	buf = make([]byte, l)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}
	name := string(buf)
	log.Println(name)
	f := &file{conn: conn, name: name, c: make(chan struct{})}
	mutex.Lock()
	filemap[rint] = f
	mutex.Unlock()
	f.wait()
	mutex.Lock()
	delete(filemap, rint)
	mutex.Unlock()
	conn.Close()
	return nil
}

func server() {
	log.Println("server")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	go RunHttp()
	go RunProxy()
	<-sc
	log.Println("catch signal, exit")
}

func client() {
	filepath := os.Args[2]
	conn, err := net.Dial("tcp", "api.iotlabs.xyz:1024")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// read rand uint32
	buf := make([]byte, 4)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}
	// print randstr share key
	randstr := strconv.FormatInt(int64(binary.BigEndian.Uint32(buf)), 36)
	log.Println("share with url: ", "http://106.15.192.203:7070/"+randstr)
	// normalize name
	if filepath[0] != '/' {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		filepath = path.Clean(dir + "/" + filepath)
	}
	filename := filepath[strings.LastIndex(filepath, "/")+1:]
	binary.BigEndian.PutUint32(buf, uint32(len(filename)))
	// send filename
	buf = append(buf, []byte(filename)...)
	_, err = conn.Write(buf)
	if err != nil {
		panic(err)
	}
	// wait send singal
	buf = make([]byte, 1)
	_, err = conn.Read(buf)
	if err != nil {
		panic(err)
	}
	// send file as soon as possible
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(conn, f)
}

func main() {
	var s, c bool
	flag.BoolVar(&s, "server", false, "--server")
	flag.BoolVar(&c, "client", false, "--client")
	flag.Parse()

	if s {
		server()
		return
	}
	if c {
		client()
		return
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
