package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"gorm.io/gorm"

	"BE-gRPC/protobuf/database"
	pb "BE-gRPC/protobuf/golang_protobuff_users"

	"google.golang.org/grpc"
)

type User struct {
	gorm.Model
	Diamonds_totals string `json:"diamonds_totals"`
	Fullname       string	`json:"fullname"`
	Username       string	`json:"username"`
	Email          string	`json:"email"`
}

type UserServer struct {
	DB *gorm.DB
}

// Update implements golang_protobuff_users.UsersServiceServer.

func (s *UserServer) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	id := req.GetId()
	diamonds_totals := req.GetDiamondsTotals()
	fullname := req.GetFullname()
	username := req.GetUsername()
	email := req.GetEmail()

	user := User{
		Model: gorm.Model{
			ID: uint(id),
		},
		Diamonds_totals: diamonds_totals,
		Fullname:       fullname,
		Username:       username,
		Email:          email,
	}

	res := s.DB.Save(&user)
	if res.Error != nil {
		return &pb.UpdateUserResponse{Success: false}, res.Error
	}

	return &pb.UpdateUserResponse{Success: true}, nil
}

func main() {
	database.DatabaseConnection()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()

	userServer := &UserServer{DB: database.DB}
	pb.RegisterUsersServiceServer(server, userServer)

	fmt.Println("Starting server on port :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
