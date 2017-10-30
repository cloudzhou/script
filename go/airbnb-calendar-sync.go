package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

type calendar struct {
	dateFrom  string
	dateTo    string
	nights    int
	available bool
}

func main() {
	log.Println("running airbnb calendar sync...")
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	airbnbCalendarQuery("1298200", 10, 2017, 1)

	go runHttp()

	<-signalChan
	log.Println("airbnb calendar sync catch signal, exit")
}

func airbnbCalendar(w http.ResponseWriter, r *http.Request) {
	airbnbCalendarQuery("1298200", 10, 2017, 1)
	http.Error(w, "ok", 200)
}

// curl https://www.airbnbchina.cn/calendar/ical/21353178.ics?s=0ac2713eacf41dbb7b5fa6310270de06

const calendarUrlf = "https://www.airbnbchina.cn/api/v2/calendar_months?key=d306zoyjsyarp7ifhu67rjxn52tv0t20&currency=CNY&locale=zh&listing_id=%s&month=%d&year=%d&count=%d"

// curl 'https://www.airbnbchina.cn/api/v2/calendar_months?key=d306zoyjsyarp7ifhu67rjxn52tv0t20&currency=CNY&locale=zh&listing_id=1298200&month=10&year=2017&count=1'
func airbnbCalendarQuery(listingId string, month int, year int, count int) []*calendar {
	url := fmt.Sprintf(calendarUrlf, listingId, month, year, count)
	status, body, err := doHttpGet(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(status, string(body))
	return nil
}

func doHttpGet(url string) (int, []byte, error) {
	log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Referer", "https://www.airbnbchina.cn/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	if err != nil {
		log.Println("1", err)
		return 500, nil, err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("2", err)
		return 500, nil, err
	}

	defer resp.Body.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, io.LimitReader(resp.Body, 1024*64))
	if err != nil {
		log.Println("3", err)
		return 500, nil, err
	}

	return resp.StatusCode, buf.Bytes(), nil
}

func runHttp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/airbnb/calendar", airbnbCalendar)

	addr := fmt.Sprintf(":8080")
	s := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1024 * 1024,
	}
	log.Println(s.ListenAndServe())
}
