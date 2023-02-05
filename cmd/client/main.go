package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IkehAkinyemi/grpc-service/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var create bool
	var read bool
	var update bool
	var delete bool

	flag.BoolVar(&create, "c", false, "Create a user request")
	flag.BoolVar(&read, "r", false, "Retrieve a user request")
	flag.BoolVar(&update, "u", false, "Update a user request")
	flag.BoolVar(&delete, "d", false, "Delete a user request")

	flag.Parse()

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := gen.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if len(os.Args) > 3 {
		fmt.Println(`Make one request at a time.

Usage:`)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "-%s \t%v\n", f.Name, f.Usage)
		})
		os.Exit(1)
	}

	userID := string(flag.Arg(0))

	if create {
		userID := createUser(ctx, client) // Create User
		log.Printf("Created User with ID: %s", userID)
	} else if read {
		newUser := readUser(ctx, client, userID) // Read User
		log.Printf("Read User: %+v", newUser)
	} else if update {
		updatedUser := updateUser(ctx, client, userID) // Update User
		log.Printf("Updated User %+v with ID: %s", updatedUser, userID)
	} else if delete {
		deleteUser(ctx, client, userID) // Delete User
	}
}
