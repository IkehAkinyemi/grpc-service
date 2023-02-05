package main

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/IkehAkinyemi/grpc-service/gen"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates random integer between min and max.
func randomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// RandomString generates a random string of length n.
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

type user struct {
	Name string
	Age  int32
}

// generateUser return a random user.
func generateUser() user {
	return user{
		Name: randomString(12),
		Age:  randomInt(1, 100),
	}
}

// createUser inserts a random user into the in-memory repository.
func createUser(ctx context.Context, client gen.UserServiceClient) string {
	user := generateUser()
	resp, err := client.CreateUser(ctx, &gen.CreateUserRequest{
		Name: user.Name,
		Age:  int32(user.Age),
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	return resp.Id
}

// readUser returns a user with the id parameter.
func readUser(ctx context.Context, client gen.UserServiceClient, id string) user {
	readResp, err := client.ReadUser(ctx, &gen.ReadUserRequest{Id: id})
	if err != nil {
		log.Fatalf("could not read user: %v", err)
	}

	return user{
		Name: readResp.Name,
		Age:  readResp.Age,
	}
}

// updateUser return an updated user info.
func updateUser(ctx context.Context, client gen.UserServiceClient, id string) user {
	_, err := client.UpdateUser(ctx, &gen.UpdateUserRequest{
		Id:  id,
		Age: randomInt(1, 100),
	})
	if err != nil {
		log.Fatalf("could not update user: %v", err)
	}

	return readUser(ctx, client, id)
}

// deleteUser deletes a user with the id parameter.
func deleteUser(ctx context.Context, client gen.UserServiceClient, id string) {
	_, err := client.DeleteUser(ctx, &gen.DeleteUserRequest{Id: id})
	if err != nil {
		log.Fatalf("could not delete user: %v", err)
	}
	log.Printf("Deleted User with ID: %s", id)
}
