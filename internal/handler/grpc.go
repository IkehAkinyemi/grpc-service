package handler

import (
	"context"

	"github.com/IkehAkinyemi/grpc-service/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	gen.UnimplementedUserServiceServer
	users map[string]user
}

type user struct {
	Name string
	Age  int32
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	id := uuid.New().String()
	h.users[id] = user{
			Name: req.Name,
			Age:  req.Age,
	}
	return &gen.CreateUserResponse{Id: id}, nil
}

func (h *Handler) ReadUser(ctx context.Context, req *gen.ReadUserRequest) (*gen.ReadUserResponse, error) {
	u, ok := h.users[req.Id]
	if !ok {
			return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return &gen.ReadUserResponse{
			Name: u.Name,
			Age:  u.Age,
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.UpdateUserResponse, error) {
	u, ok := h.users[req.Id]
	if !ok {
			return nil, status.Errorf(codes.NotFound, "user not found")
	}
	u.Name = req.Name
	u.Age = req.Age
	h.users[req.Id] = u
	return &gen.UpdateUserResponse{}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*gen.DeleteUserResponse, error) {
	delete(h.users, req.Id)
	return &gen.DeleteUserResponse{}, nil
}
