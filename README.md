## Example Blockchain with gRPC 

this project is a simple example of how `blockchain` works. in this case using the `gRPC` call procedure is because `gRPC` uses `HTTP 2.0`, which makes the data exchange process faster than using the `API call` procedure in general.

### How to Run :
1. run `go mod tidy` first to get needed package.  
2. `go run server/main.go` : to running server.
3. And on the client side , type `go run client/main.go --add` to add some blocks of hash
4. And type `go run client/main.go --list` to get block of hash .

### Credits 
> 1. [https://golang.org](https://golang.org) - The Go programming language
> 2. [gRPC](https://grpc.io) -  A high-performance, open source universal RPC framework
> 3. [Protocol Buffer](https://developers.google.com/protocol-buffers) - Protocol buffers are a language-neutral, platform-neutral extensible mechanism for serializing structured data.
> 4. [Go - Protobuf](https://github.com/golang/protobuf) - Go support for Google's protocol buffers
