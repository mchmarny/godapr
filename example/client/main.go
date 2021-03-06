package main

import (
	"encoding/json"
	"fmt"

	dapr "github.com/mchmarny/godapr/v1"
	"go.opencensus.io/trace"
)

func main() {
	// just for this demo
	ctx := trace.SpanContext{}
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
	resp, err := client.InvokeServiceWithIdentity(ctx, "example-service", "echo", msgIn)
	if err != nil {
		panic(err)
	}
	msgOut := EchoMessage{}
	if err := json.Unmarshal(resp, &msgOut); err != nil {
		panic(err)
	}
	fmt.Printf("send: ping, got: %s\n", msgOut.Response)

	// invoke output binding named 'example-http-binding'
	_, err = client.InvokeBinding(ctx, "example-http-binding", "create")
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
