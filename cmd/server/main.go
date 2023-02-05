package main

import (
	"fmt"
	"log"
	"net"

	"github.com/IkehAkinyemi/grpc-service/gen"
	"github.com/IkehAkinyemi/grpc-service/internal/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Starting the sever")
	port := 50051

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	h := handler.New()
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterUserServiceServer(srv, h)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
