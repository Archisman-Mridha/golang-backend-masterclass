package server

import (
	"context"

	"simple_bank/generated/proto"
	"simple_bank/handlers"
)

func(server *SimpleBankGRPCServer) CreateUser(
	ctx context.Context, request *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {

	return handlers.CreateUserHandler(request)
}