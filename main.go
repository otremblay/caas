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
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	err := http.ListenAndServe(":"+port, http.HandlerFunc(getTimeZone))
	if err != nil {
		log.Fatalln(err)
	}
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
