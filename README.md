# gRPC CRUD Service in Go
This repository contains an example of a simple gRPC CRUD service implemented in Go. The service allows users to create, read, update, and delete users. The service uses Protocol Buffers as the interface definition language (IDL) for defining the service API and for serializing and transmitting the data.

## Prerequisites
Go 1.14 or higher
Protocol Buffers 3.x
gRPC 1.33 or higher

## Setting up the development environment
Installing the Protocol Buffers and gRPC library for Go project: To get started, you need to install the Protocol Buffers compiler (protoc) and the gRPC library for Go. You can do this by running the following command:

```sh
# install Protocol Buffers compiler
brew install protobuf

# install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Generate the gRPC Golang code
Using the protoc compiler, compile this [project-defined Protocol Buffers message](./api/user.proto) for the user service into Golang code:

```sh
protoc --proto_path=. --go_out=. \
  --go-grpc_out=. \
  ./api/*.proto
```

Check the `/gen` for the generated Golang code.

## Running and building
Use the `go build` command to spin up the gRPC server like below:

```sh
# build server
go build -o=./grpc-server./cmd/server/*

# Run the executable
./grpc-server
```

Using the CLI client, we can either create, read, update or delete a user from the in-memory repository.

```sh
# Create a random user
go run ./cmd/client -c

# Retrieve a user using the return id from '-c' command
go run ./cmd/client -r <user_id>

# Update a user info
go run ./cmd/client -u <user_id>

# Delete a user
go run ./cmd/client -d <user_id>
```


## Testing
gRPC-service comes with mixed bag of unit and integration tests (coming soon).
To run them, use the default `go test` command:
```sh
go test ./...
```

## Security

If you discover a security vulnerability within gRPC-service, please raise an issue on this repository or send an e-mail to **mrikehchukwuka@gmail.com**.

All issues and reports will be promptly addressed, and you'll be credited accordingly.

## Future work & contribution
You could help continuing its development by:

- Contribute to the source code
- Suggest new features and report issues
