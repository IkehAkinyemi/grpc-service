# gRPC CRUD Service in Go
This repository contains an example of a simple gRPC CRUD service implemented in Go. The service allows users to create, read, update, and delete users. The service uses Protocol Buffers as the interface definition language (IDL) for defining the service API and for serializing and transmitting the data.

# Prerequisites
Go 1.14 or higher
Protocol Buffers 3.x
gRPC 1.33 or higher

# Setting up the development environment
Installing the Protocol Buffers and gRPC library for Go project: To get started, you need to install the Protocol Buffers compiler (protoc) and the gRPC library for Go. You can do this by running the following command:

```
# install Protocol Buffers compiler
brew install protobuf

# install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Generate the gRPC Golang code
Using the protoc compiler, compile this [project-defined Protocol Buffers message](./api/user.proto) for the user service into Golang code:

```
protoc --proto_path=. --go_out=. \
  --go-grpc_out=. \
  ./api/*.proto
```

Check the `/gen` for the generated Golang code.