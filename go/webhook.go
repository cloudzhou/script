package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	data := make([]byte, 1024*1024)
	n, err := io.ReadFull(r.Body, data)
	if err != nil && err != io.ErrUnexpectedEOF {
		fmt.Println(err)
		io.WriteString(w, "failed(read)\n")
		return
	}
	data = data[0:n]
	obj := make(map[string]interface{})
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, "failed(json)\n")
		return
	}
	ref := obj["ref"].(string)
	after := obj["after"].(string)
	fmt.Println("ref: %s, commit: %s", ref, after)
	generateBin(ref, after)
	io.WriteString(w, "ok")
}

func getVersionLifecycle(ref string) (string, string) {
	ss := strings.Split(ref, "/")
	name := ss[len(ss)-1]
	ss = strings.Split(name, "-")
	if len(ss) != 2 {
		return "", ""
	}
	return ss[0], ss[1]
}

func generateBin(ref string, commit string) {
	version, lifecycle := getVersionLifecycle(ref)
	if version == "" || lifecycle == "" {
		return
	}
	fmt.Println("generate bin with version: %s, lifecycle: %s, commit: %s", version, lifecycle, commit)
	url := "http://iot.espressif.cn/v1/products/1/roms/"
	body := []byte(fmt.Sprintf(`{"productRoms":[{"version":"%s","life_cycle":"%s","recommended":1,"files":[{"path":"6f30accbc94a19224eb71f2e392ed32d.bin","name":"user1.bin"}, {"path":"6f30accbc94a19224eb71f2e392ed32d.bin","name":"user2.bin"}]}]}`, version, lifecycle))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "token xxxxxxxxxxxxxxxxxxxxxxxxxx")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8444", nil)
}
