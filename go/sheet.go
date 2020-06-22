package main

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	domain  = "http://localhost:8080"
	datadir = "/tmp" // .csv .json .secret
	prefix  = "/sheet"
	maxlen  = 1 << 20
)

type handler struct{}

func (h handler) getAuthorization(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	ss := strings.SplitN(authorization, " ", 2)
	return ss[1]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if !strings.HasPrefix(path, prefix) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var name, suffix string
	trim := strings.TrimPrefix(path, prefix)
	if strings.HasPrefix(trim, "/") {
		trim = trim[1:]
	}
	name = trim
	if index := strings.LastIndex(trim, "."); index > 0 {
		name = trim[0:index]
		suffix = trim[index+1:]
	}

	var err error
	switch r.Method {
	case http.MethodGet:
		err = h.get(w, r, name, suffix)
	case http.MethodPost:
		err = h.post(w, r, name)
	case http.MethodDelete:
		err = h.delete(w, r, name)
	}

	if err != nil {
		log.Println(err)
		if os.IsNotExist(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h handler) readfile(path string) (content []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	content = make([]byte, fileInfo.Size())
	_, err = io.ReadFull(file, content)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return
}

// get /sheet/name{.html/.csv/.json}
func (h handler) get(w http.ResponseWriter, r *http.Request, name, suffix string) (err error) {
	switch suffix {
	case "html":
		w.Header().Add("Content-type", "text/html; charset=UTF-8")
		url := fmt.Sprintf("%s%s/%s.json", domain, prefix, name)
		_, err = w.Write([]byte(fmt.Sprintf(indexfmt, url)))
	case "csv":
		w.Header().Add("Content-type", "text/csv; charset=UTF-8")
		w.Header().Add("Content-Disposition", "attachment; filename="+name+".csv")
	case "json":
		w.Header().Add("Content-type", "application/json")
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if suffix == "csv" || suffix == "json" {
		namebase64 := base64.StdEncoding.EncodeToString([]byte(name))
		filename := namebase64 + "." + suffix
		path := filepath.Join(datadir, filename)
		file, err := os.Open(path)
		log.Println(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(w, file)
	}

	return
}

// post /sheet[/name]
func (h handler) post(w http.ResponseWriter, r *http.Request, name string) error {
	if name == "" {
		b := make([]byte, 4)
		_, err := rand.Read(b)
		if err != nil {
			return err
		}
		name = hex.EncodeToString(b)
	}

	// body
	b, err := ioutil.ReadAll(io.LimitReader(r.Body, maxlen))
	if err != nil {
		return err
	}

	namebase64 := base64.StdEncoding.EncodeToString([]byte(name))

	// authorization
	authpath := filepath.Join(datadir, namebase64+".authorization")
	if _, err := os.Stat(authpath); os.IsNotExist(err) {
		// if not exist, create authorization
		csvfile, err := os.Create(authpath)
		if err != nil {
			return err
		}
		defer csvfile.Close()
		_, err = csvfile.Write([]byte(h.getAuthorization(r)))
		if err != nil {
			return err
		}
	} else {
		// if exist, check authorization
		authorization, err := h.readfile(authpath)
		if err != nil {
			return err
		}
		authorizationHeader := h.getAuthorization(r)
		if string(authorization) != authorizationHeader {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return nil
		}
	}

	// csv
	cvspath := filepath.Join(datadir, namebase64+".csv")
	csvfile, err := os.Create(cvspath)
	if err != nil {
		return err
	}
	defer csvfile.Close()
	var arrays [][]interface{}
	err = json.Unmarshal(b, &arrays)
	if err != nil {
		return err
	}
	var records [][]string
	for _, array := range arrays {
		var record []string
		for _, item := range array {
			record = append(record, fmt.Sprintf("%v", item))
		}
		records = append(records, record)
	}
	err = csv.NewWriter(csvfile).WriteAll(records)
	if err != nil {
		return err
	}

	// json
	jsonpath := filepath.Join(datadir, namebase64+".json")
	jsonfile, err := os.Create(jsonpath)
	if err != nil {
		return err
	}
	defer jsonfile.Close()
	_, err = jsonfile.Write(b)
	if err != nil {
		return err
	}

	// output new link
	_, err = w.Write([]byte(fmt.Sprintf("%s%s/%s.html\n", domain, prefix, name)))
	return err
}

// delete /sheet/name
func (h handler) delete(w http.ResponseWriter, r *http.Request, name string) error {
	namebase64 := base64.StdEncoding.EncodeToString([]byte(name))
	filename := namebase64 + ".authorization"
	path := filepath.Join(datadir, filename)

	authorization, err := h.readfile(path)
	if err != nil {
		return err
	}
	authorizationHeader := h.getAuthorization(r)
	if string(authorization) != authorizationHeader {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return nil
	}

	for _, suffix := range []string{".json", ".csv", ".authorization"} {
		err := os.Remove(filepath.Join(datadir, namebase64+suffix))
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

const indexfmt = `<!DOCTYPE html>
<html>
<head>
	<link rel='stylesheet' type='text/css' href='https://gitshell.com/static/sheet/spreadsheetStyle.css'/>
	<script src='https://gitshell.com/static/sheet/spreadsheetStyle.js'></script>
	<script>
		window.onload = function() {
			var xhttp = new XMLHttpRequest();
			xhttp.onreadystatechange = function() {
				if (this.readyState == 4 && this.status == 200) {
					var container = document.getElementById("container");
					table = new Spreadsheet({
						context: container,
						data: JSON.parse(this.responseText)
					})
				}
			};
			xhttp.open("GET", "%s", true);
			xhttp.send();
		}
	</script>
</head>

<body>
	<div name='table' id='container'></div><br/>
</body>
</html>
`

func init() {
	rand.Seed(time.Now().UnixNano())
}
