package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiUrl       = "https://gitlab.espressif.cn:6688/api/v3"
	token        = "espressif-gitlab-bot"
	privateToken = "zStdT1TytGdNEzyTyaZT"
)

func main() {
	log.Println("start espressif bot...")
	httpHandler := NewHttpHandler()
	s := &http.Server{
		Addr:           ":5001",
		Handler:        httpHandler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 5,
	}
	s.ListenAndServe()
	log.Println("espressif bot exit...")
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

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/bot", "/bot/":
		h.doBot(w, req)
	default:
		http.Error(w, "Path Not Found", 404)
		return
	}
}

func (h *HttpHandler) doBot(w http.ResponseWriter, req *http.Request) {
	if token != req.Header.Get("X-Gitlab-Token") {
		http.Error(w, "Require X-Gitlab-Token", 403)
		return
	}

	post, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Read Failed", 500)
		return
	}
	defer req.Body.Close()
	body := make(map[string]interface{})
	err = json.Unmarshal(post, &body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal Json Failed", 400)
		return
	}
	objectKind := h.GetValue("object_kind", body).(string)
	switch objectKind {
	case "note":
		h.doNote(w, req, body)
	default:
		http.Error(w, "Get Empty object_kind", 400)
		return
	}
}

func (h *HttpHandler) doNote(w http.ResponseWriter, req *http.Request, body map[string]interface{}) {
	noteableType := h.GetValue("object_attributes.noteable_type", body).(string)
	switch noteableType {
	case "MergeRequest":
		projectId := int(h.GetValue("merge_request.source_project_id", body).(float64))
		mergeId := int(h.GetValue("merge_request.id", body).(float64))
		uri := fmt.Sprintf("%s/projects/%d/merge_requests/%d/notes?private_token=%s", apiUrl, projectId, mergeId, privateToken)
		h.Post(uri, "body", "Triggered unit tests run")
	default:
		http.Error(w, "Bad object_attributes.noteable_type", 400)
		return
	}
}

func (h *HttpHandler) Post(uri string, key string, value string) error {
	data := make(url.Values)
	data[key] = []string{value}
	resp, err := http.PostForm(uri, data)
	log.Println(resp)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (h *HttpHandler) GetValue(arg string, m map[string]interface{}) interface{} {
	if m == nil {
		return nil
	}

	var v interface{}
	var ok bool

	s := strings.Split(arg, ".")
	for i, k := range s {
		v, ok = m[k]
		if !ok {
			return nil
		}
		if i == len(s)-1 {
			return v
		}
		switch vv := v.(type) {
		case map[string]interface{}:
			m = vv
		default:
			return nil
		}
	}
	return nil
}
