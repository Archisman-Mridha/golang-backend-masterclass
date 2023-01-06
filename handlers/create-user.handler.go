package handlers

import (
	"simple_bank/generated/proto"
	"simple_bank/middlewares"
)

func CreateUserHandler(request *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {

	// pass request through middleware
	middlewares.CreateUserMiddleware(request)

	return &proto.CreateUserResponse{ }, nil
}