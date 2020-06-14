# godapr client example 

## Service 
Inside of the [example/service](./example/service) directory run this:

```shell
dapr run --app-id example-service \
         --app-port 8080 \
         --protocol http \
         --components-path ./components \
         go run main.go 
```

## Client 
Inside of the [example/client](./example/client) directory run this:

```shell
dapr run --app-id example-client \
         --components-path ./components \
         go run main.go
```

## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.

## License
This software is released under the [Apache v2 License](./LICENSE)
