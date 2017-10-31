package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
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
	calendars := airbnbCalendarQuery("1298200", 10, 2017, 1)
	var buf bytes.Buffer
	buf.WriteString(begin)
	for _, calendar := range calendars {
		vevent := fmt.Sprintf(veventf, calendar.dateTo, calendar.dateFrom, calendar.dateFrom, calendar.dateTo, calendar.nights)
		buf.WriteString(vevent)
	}
	buf.WriteString(end)
	http.Error(w, buf.String(), 200)
}

const begin = `BEGIN:VCALENDAR
PRODID;X-RICAL-TZSOURCE=TZINFO:-//Airbnb Inc//Hosting Calendar 0.8.8//EN
CALSCALE:GREGORIAN
VERSION:2.0
`
const veventf = `BEGIN:VEVENT
DTEND;VALUE=DATE:%s
DTSTART;VALUE=DATE:%s
UID:-orcx5rvky2gs--ep729h7uh1t6@airbnb.com
DESCRIPTION:CHECKIN: %s\nCHECKOUT: %s\nNIGHTS: %d\nPHONE:
 13800138000\nEMAIL: hi@guest.airbnb.com\nPROP
 ERTY: 【江景远眺】三室两厅两卫 19层电梯复式 景观房\n
SUMMARY:weiwei
LOCATION:【江景远眺】三室两厅两卫 19层电梯复式 景观房
END:VEVENT
`
const end = `END:VCALENDAR`

// curl https://www.airbnbchina.cn/calendar/ical/21353178.ics?s=0ac2713eacf41dbb7b5fa6310270de06

const calendarUrlf = "https://www.airbnbchina.cn/api/v2/calendar_months?key=d306zoyjsyarp7ifhu67rjxn52tv0t20&currency=CNY&locale=zh&listing_id=%s&month=%d&year=%d&count=%d"

// curl 'https://www.airbnbchina.cn/api/v2/calendar_months?key=d306zoyjsyarp7ifhu67rjxn52tv0t20&currency=CNY&locale=zh&listing_id=1298200&month=10&year=2017&count=1'
func airbnbCalendarQuery(listingId string, month int, year int, count int) []*calendar {
	url := fmt.Sprintf(calendarUrlf, listingId, month, year, count)
	log.Println(url)
	// status, body, err := doHttpGet(url)
	var err error
	status, body, err := 200, jj, nil // doHttpGet(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println(status, string(body))
	calendars := calendarUnavailable(string(body))
	return calendars
}

func calendarUnavailable(body string) []*calendar {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(body), &m)
	if err != nil {
		log.Println(err)
		return nil
	}

	calendarMonths, ok := getValue(m, "calendar_months").([]interface{})
	if !ok {
		log.Println("can not found calendar_months")
		return nil
	}

	calendars := []*calendar{}
	dateFrom := ""
	dateTo := ""
	nights := 0
	preAvailable := true
	for _, x := range calendarMonths {
		calendarMonth, ok := x.(map[string]interface{})
		if !ok {
			log.Println("can not found calendar_month")
			return nil
		}
		days, ok := getValue(calendarMonth, "days").([]interface{})
		if !ok {
			log.Println("can not found days")
			return nil
		}
		for _, y := range days {
			day, ok := y.(map[string]interface{})
			if !ok {
				log.Println("can not found day")
			}
			available, ok1 := getValue(day, "available").(bool)
			date, ok2 := getValue(day, "date").(string)
			if !ok1 || !ok2 {
				log.Println("available and date not ok")
			}
			if !available {
				if preAvailable {
					dateFrom = date
					preAvailable = false
					nights = 1
				} else {
					nights += 1
				}
			}
			dateTo = date
			if available && !preAvailable {
				preAvailable = true
				calendars = append(calendars, &calendar{dateFrom: dateFrom, dateTo: dateTo, nights: nights, available: false})
				log.Println(dateFrom, dateTo, nights)
				dateFrom, dateTo, nights = "", "", 0
			}
			log.Println(available, date)
		}
	}
	d, err := time.Parse("2006-01-02", dateTo)
	if err != nil {
		log.Println(err)
		return nil
	}
	d = d.Add(24 * time.Hour)
	dateTo = d.Format("2006-01-02")

	if nights > 0 {
		calendars = append(calendars, &calendar{dateFrom: dateFrom, dateTo: dateTo, nights: nights, available: false})
		log.Println(dateFrom, dateTo, nights)
	}
	return calendars
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

func queryInt(r *http.Request, arg string) int {
	value := r.URL.Query().Get(arg)
	if value == "" {
		return 0
	}
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Println(err)
		return 0
	}
	return int(i)
}

func queryString(r *http.Request, arg string) string {
	return r.URL.Query().Get(arg)
}

func getValue(m map[string]interface{}, arg string) interface{} {
	if m == nil {
		return nil
	}

	var v interface{}
	var ok bool
	var index int64

	s := strings.Split(arg, ".")
	for ii, k := range s {
		i := strings.Index(k, "[")
		if i > 0 {
			j := strings.Index(k, "]")
			if j > 0 && j > i {
				index, _ = strconv.ParseInt(k[i+1:j], 10, 64)
				k = k[0:i]
			}

		}
		v, ok = m[k]
		if !ok {
			return nil
		}
		if i < 0 && ii == len(s)-1 {
			return v
		}
		switch vv := v.(type) {
		case map[string]interface{}:
			m = vv
		case []interface{}:
			if ii == len(s)-1 {
				return vv[int(index)]
			} else {
				m = vv[int(index)].(map[string]interface{})
			}
		default:
			return nil
		}
	}
	return nil
}

const jj = `{"calendar_months":[{"abbr_name":"10月","day_names":["一","二","三","四","五","六","日"],"days":[{"available":false,"date":"2017-09-25"},{"available":false,"date":"2017-09-26"},{"available":false,"date":"2017-09-27"},{"available":true,"date":"2017-09-28"},{"available":false,"date":"2017-09-29"},{"available":false,"date":"2017-09-30"},{"available":false,"date":"2017-10-01"},{"available":false,"date":"2017-10-02"},{"available":false,"date":"2017-10-03"},{"available":false,"date":"2017-10-04"},{"available":false,"date":"2017-10-05"},{"available":false,"date":"2017-10-06"},{"available":false,"date":"2017-10-07"},{"available":false,"date":"2017-10-08"},{"available":true,"date":"2017-10-09"},{"available":false,"date":"2017-10-10"},{"available":false,"date":"2017-10-11"},{"available":false,"date":"2017-10-12"},{"available":false,"date":"2017-10-13"},{"available":false,"date":"2017-10-14"},{"available":true,"date":"2017-10-15"},{"available":true,"date":"2017-10-16"},{"available":true,"date":"2017-10-17"},{"available":false,"date":"2017-10-18"},{"available":false,"date":"2017-10-19"},{"available":false,"date":"2017-10-20"},{"available":false,"date":"2017-10-21"},{"available":false,"date":"2017-10-22"},{"available":false,"date":"2017-10-23"},{"available":false,"date":"2017-10-24"},{"available":false,"date":"2017-10-25"},{"available":false,"date":"2017-10-26"},{"available":true,"date":"2017-10-27"},{"available":false,"date":"2017-10-28"},{"available":false,"date":"2017-10-29"},{"available":false,"date":"2017-10-30"},{"available":false,"date":"2017-10-31"},{"available":false,"date":"2017-11-01"},{"available":true,"date":"2017-11-02"},{"available":true,"date":"2017-11-03"},{"available":false,"date":"2017-11-04"},{"available":false,"date":"2017-11-05"}],"dynamic_pricing_updated_at":"2017-10-16T10:43:44+00:00","month":10,"name":"十月","year":2017,"listing_id":1298200}],"metadata":{}}`
