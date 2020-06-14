# godapr (simple dapr HTTP client)

Dapr has gRPC and REST APIs. For `go`, there is an auto-generated [gRPC SDK](https://github.com/dapr/go-sdk) and the developers can also implement their own HTTP calls to the REST API. When invoking the dapr REST APIs there usually is lot's of redundant code building request, parsing responses, and dealing with traces. I create this client to simplify Dapr integrations and minimize code duplication.

> Note, I submitted a [PR](https://github.com/dapr/go-sdk/pull/18) with similar enhancements to the [go-sdk](https://github.com/dapr/go-sdk) already submitted [#18](https://github.com/dapr/go-sdk/pull/18)

## Usage

To use `godapr` first get the library

```shell
go get github.com/mchmarny/godapr/v1
```

### Create Client

To use `godapr` library in your code, first import it

```go
import dapr "github.com/mchmarny/godapr/v1"
```

Then create a `godapr` client with the `dapr` server defaults

```go
client := dapr.NewClient()
```

or if you need to specify non-default dapr port

```go
client := dapr.NewClientWithURL("http://localhost:3500")
```

### State


#### Save Data

To save state using the the "reasonable" defaults:

```go
state := "my data"
err := client.SaveState(ctx, "store-name", "id-123", state)
```

You can also persist objects

```go
person := &Person{ Name: "Example John", Age: 35 }
err := client.SaveState(ctx, "store-name", "id-123", person)
```

For more control, you can also create the `StateData` object

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

#### Get Data

To get state data you can either use the client defaults ("strong" Consistency, "last-write" Concurrency)

```go
data, err := client.GetState(ctx, "store-name", "record-key")
```

Or, for more control, define your own state options

```go
opt := &StateOptions{
    Consistency: "eventual",
    Concurrency: "first-write",
}

data, err := client.GetStateWithOptions(ctx, "store-name", "record-key", opt)
```

#### Delete Data 

Similarly with deleting, you can either use the defaults

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

Similarly with binding, you can invoke binding without any data

```go
out, err := client.InvokeBinding(ctx, "binding-name", "create")
```

Or, with an instance of your own struct

```go
err := client.InvokeBindingWithIdentity(ctx, "binding-name", "create", person)
```

Or, for more control, with an instance of the `BindingData`


```go
data := &BindingData{
    Data:      []byte("your content"),
    Operation: "create",
    Metadata:  map[string]string{ "k1":"v1", "k2": "v2" },
}
err := client.InvokeBindingWithData(ctx, "binding-name", "create", data)
```

### Service Invocation 


Similarly with service to service invocation, you can invoke without any data

```go
out, err := client.InvokeService(ctx, "service-name", "method-name")
```

Or, with an instance of your own struct

```go
err := client.InvokeServiceWithIdentity(ctx, "service-name", "method-name", person)
```

Or, invoke it directly with your own content 


```go
data := []byte("your content")
err := client.InvokeServiceWithData(ctx, "binding-name", "create", data)
```

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.

## License
This software is released under the [Apache v2 License](./LICENSE)
