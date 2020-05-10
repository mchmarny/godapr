package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/bind", BindHandler)
	http.HandleFunc("/echo", EchoHandler)
	http.ListenAndServe(":8080", nil)
}

// BindHandler responds to everything with OK
func BindHandler(w http.ResponseWriter, r *http.Request) {
	m := &BindMessage{
		Message: "OK",
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		handleError(w, err)
	}
}

// EchoHandler responds to a message
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	m := EchoMessage{}
	if err := json.Unmarshal(b, &m); err != nil {
		handleError(w, err)
	}
	m.Response = "pong"
	if err := json.NewEncoder(w).Encode(m); err != nil {
		handleError(w, err)
	}
}

func handleError(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}

// BindMessage represents the simplest message
type BindMessage struct {
	Message string `json:"msg"`
}

// EchoMessage holds the request and response
type EchoMessage struct {
	Request  string `json:"req"`
	Response string `json:"res"`
}
