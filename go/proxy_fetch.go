package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ProxyHandler struct {
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/proxy":
		values := req.URL.Query()
		urlstr := values.Get("url")
		switch urlstr {
		case "/youtube/v3/search":
			p.youtubeSearch(w, req)
		default:
			p.proxy(w, req)
		}
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func (p *ProxyHandler) get(urlstr string) ([]byte, int, int, http.Header, error) {
	if urlstr == "" {
		return nil, 0, 400, nil, errors.New("parameter url not found")
	}
	url, err := url.Parse(urlstr)
	if err != nil {
		return nil, 0, 500, nil, errors.New(fmt.Sprint("url parse error: %s", err))
	}
	host, path := url.Host, url.Path
	fmt.Println(host, path)
	client := http.Client{
		Timeout: time.Duration(3 * time.Second),
	}
	resp, err := client.Get(urlstr)
	if err != nil || resp.StatusCode != 200 {
		return nil, 0, 500, nil, errors.New("proxy url return http status != 200")
	}
	body := make([]byte, 1024*1024)
	defer resp.Body.Close()
	n, err := io.ReadFull(resp.Body, body)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, 0, 500, nil, errors.New("proxy body content length great then 1024 * 1024")
	}
	return body, n, 200, resp.Header, nil
}

func (p *ProxyHandler) proxy(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	urlstr := values.Get("url")
	body, n, httpStatus, proxyHeader, err := p.get(urlstr)
	if err != nil {
		http.Error(w, err.Error(), httpStatus)
		return
	}
	w.Header().Set("Content-Type", proxyHeader.Get("Content-Type"))
	w.Header().Set("Content-Length", strconv.Itoa(n))
	w.WriteHeader(200)
	w.Write(body[0:n])
}

func (p *ProxyHandler) youtubeSearch(w http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()
	q := values.Get("q")
	callback := values.Get("callback")
	order := values.Get("order")
	if q == "" || callback == "" || order == "" {
		http.Error(w, "where are parameters: q && callback && order", 500)
		return
	}
	urlstr := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?key=AIzaSyCeWIOjxc80TqQ-ah1mNj2krnr3fTrcF-U&part=snippet&q=%s&type=video&maxResults=10&callback=%s&order=%s", q, callback, order)
	body, n, httpStatus, proxyHeader, err := p.get(urlstr)
	if err != nil {
		http.Error(w, err.Error(), httpStatus)
		return
	}
	w.Header().Set("Content-Type", proxyHeader.Get("Content-Type"))
	w.Header().Set("Content-Length", strconv.Itoa(n))
	w.WriteHeader(200)
	w.Write(body[0:n])
}

func main() {
	proxyHandler := &ProxyHandler{}
	s := &http.Server{
		Addr:           ":8080",
		Handler:        proxyHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 10,
	}
	log.Fatal(s.ListenAndServe())
}
