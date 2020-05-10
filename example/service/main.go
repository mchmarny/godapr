package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", EchoHandler)
	http.ListenAndServe(":8080", nil)
}

// EchoHandler responds to a message
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	m := EchoMessage{}
	if err := json.Unmarshal(b, &m); err != nil {
		panic(err)
	}
	m.Response = "pong"
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

// EchoMessage holds the request and response
type EchoMessage struct {
	Request  string `json:"req"`
	Response string `json:"res"`
}
