package main

import (
	"context"
	"encoding/json"
	"fmt"

	dapr "github.com/mchmarny/godapr"
)

func main() {
	// just for this demo
	ctx := context.Background()
	data := &Person{
		Name: "Test Person",
		Age:  40,
	}

	// create the client
	client := dapr.NewClient()

	// publish a message to the topic messagebus
	err := client.Publish(ctx, "messagebus", data)
	if err != nil {
		panic(err)
	}
	fmt.Println("data published")

	// save state with the key key1
	err = client.SaveState(ctx, "statestore", "key1", data)
	if err != nil {
		panic(err)
	}
	fmt.Println("data saved")

	// get state for key key1
	dataOut, err := client.GetState(ctx, "statestore", "key1")
	if err != nil {
		panic(err)
	}
	data2 := Person{}
	if err := json.Unmarshal(dataOut, &data2); err != nil {
		panic(err)
	}
	if data2.Name != data.Name {
		panic("invalid data retrieved from store")
	}
	fmt.Println("data got")

	// delete state for key key1
	err = client.DeleteState(ctx, "statestore", "key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("data deleted")

	// invoke a method called MyMethod on another dapr enabled service with id client
	msgIn := &EchoMessage{Request: "ping"}
	resp, err := client.InvokeService(ctx, "example-service", "echo", msgIn)
	if err != nil {
		panic(err)
	}
	msgOut := EchoMessage{}
	if err := json.Unmarshal(resp, &msgOut); err != nil {
		panic(err)
	}
	fmt.Printf("Request: ping, Response: %s\n", msgOut.Response)

	// invoke output binding named 'example-http-binding'
	// uses https://http2.pro/doc/api to check for HTTP/2
	err = client.InvokeBinding(ctx, "example-http-binding", data)
	if err != nil {
		panic(err)
	}
	fmt.Println("binding invoked")

	fmt.Println("DONE")
}

// Person is a test object for this example
type Person struct {
	Name string
	Age  int
}

// EchoMessage holds the request and response
type EchoMessage struct {
	Request  string `json:"req"`
	Response string `json:"res"`
}
