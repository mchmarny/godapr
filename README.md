# godapr (simple dapr HTTP client)

Dapr has gRPC and REST APIs. For `go`, there is the auto-generated [gRPC SDK](https://github.com/dapr/go-sdk) that covers the complete spectrum of dapr API. Developers can also implement their own HTTP calls to the REST API. When invoking the dapr REST APIs there usually is lot's of redundant code building request and parsing responses, so I create this simple Dapr client to simplify Dapr integrations and minimize code duplication.

## Usage

To use `godapr` first get the library

```shell
go get github.com/mchmarny/godapr
```

### Create Client

To use `godapr` library in your code, first import it

```go
import dapr "github.com/mchmarny/godapr"
```

Then create a `godapr` client with the `dapr` server defaults

```go
client := dapr.NewClient()
```

or if you need to specify non-default dapr port

```go
client := dapr.NewClientWithURL("http://localhost:3500")
```

> consider getting the dapr server URL from environment variable

### State

#### Get Data

To get state data you can either use the client defaults ("strong" Consistency, "last-write" Concurrency)

```go
data, err := client.GetState(ctx, "store-name", "record-key")
```

Or define your own state options

```go
opt := &StateOptions{
    Consistency: "eventual",
    Concurrency: "first-write",
}

data, err := client.GetStateWithOptions(ctx, "store-name", "record-key", opt)
```

#### Save Data

Similarly with saving state, assuming you have your own person object for example

```go
person := &Person{
    Name: "Example John",
    Age: 35,
}
```

you can either use the defaults

```go
err := client.SaveState(ctx, "store-name", "record-key", person)
```

Or define your own state data object

```go
data := &StateData{
    Key: "id-123",
    Value: person,
    Options: &StateOptions{
        Consistency: "eventual",
        Concurrency: "first-write",
    },
}

err := client.SaveStateWithData(ctx, "store-name", data)
```

#### Delete Data 


Similarly with deleting ata... you can either use the defaults


```go
err := client.DeleteState(ctx, "store-name", "record-key")
```

Or define your own state data object

```go
opt := &StateOptions{
    Consistency: "eventual",
    Concurrency: "first-write",
}

err := client.DeleteStateWithOptions(ctx, "store-name", opt)
```

### Events

To publish events to a topic you can pass instance of your own struct

```go
err := client.Publish(ctx, "topic-name", person)
```

Or send the raw content in bytes 

```go
data := []byte("hi")
err := client.PublishWithData(ctx, "topic-name", data)
```


### Binding

Similarly with binding you can use the default method to send instance of your own struct

```go
err := client.InvokeBinding(ctx, "binding-name", person)
```

Or send the raw content in bytes 

```go
data := []byte("hi")
err := client.InvokeBindingWithData(ctx, "topic-name", data)
```

### Service Invocation 

Finally, for service, you can either 

```go
out, err := client.InvokeService(ctx, "service-name", "method-name", person)
```

Or serialize the person yourself and 

```go
content, _ := json.Marshal(data)
out, err := client.InvokeServiceWithData(ctx, "service-name", "method-name", content)
```



## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.

## License
This software is released under the [Apache v2 License](./LICENSE)
