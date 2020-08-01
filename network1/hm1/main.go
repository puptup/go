package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type request struct {
	Host       string              `json:"host"`
	UserAgent  string              `json:"user_agent"`
	RequestURL string              `json:"request_uri"`
	Headers    map[string][]string `json:"headers"`
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	var req request
	req.Host = r.Host
	req.UserAgent = r.UserAgent()
	req.RequestURL = r.RequestURI
	req.Headers = r.Header

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
