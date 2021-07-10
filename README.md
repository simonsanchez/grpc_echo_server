This is a gRPC server implementation of the Echo service [here](https://github.com/simonsanchez/grpc_echo_proto_two).

# Running the service

```
go run .
```

## Calling RPCs

First install `ghz` with

```
brew install ghz
```

From the root of the **companion repo** mentioned above, run the following

### v2

```
cd v2
ghz --insecure --proto ./echo.proto --call echo.EchoService.Baz 0.0.0.0:10000 -n 10000 -c 100 --connections=50 -d '{"message": "hello"}'
```

### v0 or v1

```
# do not cd into the `v2` directory
ghz --insecure --proto ./echo.proto --call echo.EchoService.Foo 0.0.0.0:10000 -n 10000 -c 100 --connections=50 -d '{"message": "hello"}'
```
