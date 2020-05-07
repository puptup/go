package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type request struct {
	Host       string  `json:"host"`
	UserAgent  string  `json:"user_agent"`
	RequestURL string  `json:"request_uri"`
	Headers    headers `json:"headers"`
}

type headers struct {
	Accept    []string `json:"Accept"`
	UserAgent []string `json:"User-Agent"`
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	var req request
	req.Host = r.Host
	req.UserAgent = r.UserAgent()
	req.RequestURL = r.RequestURI
	req.Headers.Accept = r.Header["Accept"]
	req.Headers.UserAgent = r.Header["UserAgent"]

	js, err := json.Marshal(req)
	if err != nil {
		log.Panic(err)
	}
	w.Write(js)
}

func main() {

	http.HandleFunc("/", getRequest)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
