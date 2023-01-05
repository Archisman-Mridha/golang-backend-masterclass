package server

import (
	"context"

	"simple_bank/generated/proto"
)

func(server *SimpleBankGRPCServer) CreateUser(
	ctx context.Context, request *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {

	return &proto.CreateUserResponse{ }, nil
}