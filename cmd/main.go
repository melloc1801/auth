package main

import (
	desc "auth/pkg/user_v1"
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	var id int64 = 1

	user := req.GetUser()

	fmt.Println("Create request")
	fmt.Println(color.GreenString("", "Id: ", id))
	fmt.Println(color.GreenString("", "Name: ", user.Name))
	fmt.Println(color.GreenString("", "Email: ", user.Email))
	fmt.Println(color.GreenString("", "Role: ", user.Role))
	fmt.Println(color.GreenString("", "Password: ", user.Password))
	fmt.Println(color.GreenString("", "PasswordConfirm: ", user.PasswordConfirm))
	fmt.Println("===================================================")

	return &desc.CreateUserResponse{
		Id: id,
	}, nil
}

func (s *server) Get(_ context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	timestamp := timestamppb.New(time.Now())
	user := &desc.GetUserResponse{
		Id:        req.GetId(),
		Name:      "Root",
		Email:     "root@mail.ru",
		Role:      desc.Role_ADMIN,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}

	fmt.Println("Get request")
	fmt.Println(color.GreenString("", "Id: ", user.Id))
	fmt.Println(color.GreenString("", "Name: ", user.Name))
	fmt.Println(color.GreenString("", "Email: ", user.Email))
	fmt.Println(color.GreenString("", "Role: ", user.Role))
	fmt.Println(color.GreenString("", "CreatedAt: ", timestamp))
	fmt.Println(color.GreenString("", "UpdatedAt: ", timestamp))
	fmt.Println("===================================================")

	return user, nil
}

func (s *server) Update(_ context.Context, req *desc.UpdateUserRequest) (*empty.Empty, error) {
	fmt.Println("Update request")
	fmt.Println(color.GreenString("", "Id: ", req.Id))
	fmt.Println(color.GreenString("", "Email: ", req.Email))
	fmt.Println(color.GreenString("", "Name: ", req.Name))
	fmt.Println(color.GreenString("", "Role: ", req.Role))
	fmt.Println("===================================================")

	return &empty.Empty{}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteUserRequest) (*empty.Empty, error) {
	fmt.Println("Delete request")
	fmt.Println(color.GreenString("", "Id: ", req.Id))
	fmt.Println("===================================================")

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
