package godapr

/*

Dapr has gRPC and REST APIs. For `go`, there is the auto-generated
[gRPC SDK](https://github.com/dapr/go-sdk) that covers the complete
spectrum of dapr API. Developers can also implement their own HTTP
calls to the REST API. When invoking the dapr REST APIs there
usually is lot's of redundant code building request and parsing
responses, so I create this simple Dapr client to simplify Dapr
integrations and minimize code duplication.

Usage

	go get github.com/mchmarny/godapr/v1


Import

	import dapr "github.com/mchmarny/godapr/v1"

Client

Then create a client with the server defaults

	client := dapr.NewClient()

or if you need to specify non-default dapr port


	client := dapr.NewClientWithURL("http://localhost:3500")

Consider getting the dapr server URL from environment variable

See: https://github.com/mchmarny/godapr

Disclaimer

This is my personal project and it does not represent my
employer. I take no responsibility for issues caused by this
code. I do my best to ensure that everything works, but if
something goes wrong, my apologies is all you will get.

*/