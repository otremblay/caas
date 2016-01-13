package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: " + os.Args[0] + " <port number>")
	}
	http.ListenAndServe(":"+os.Args[1], http.HandlerFunc(getTimeZone))
}

func getTimeZone(rw http.ResponseWriter, req *http.Request) {
	timezone := strings.TrimPrefix(req.URL.Path, "/")
	location, err := time.LoadLocation(timezone)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, err.Error())
		return
	}
	if location == nil {
		rw.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(rw, "Time zone not found")
		return
	}
	t := time.Now().In(location)
	jsonEncoder := json.NewEncoder(rw)
	jsonEncoder.Encode(Time{Date: t.Format("2006/01/02"), Time: t.Format("15:04:05"), Full: t.String()})
}

type Time struct {
	Date string
	Time string
	Full string
}
