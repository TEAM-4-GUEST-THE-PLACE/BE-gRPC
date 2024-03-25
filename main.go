package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

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
	DeletedAt      gorm.DeletedAt `gorm: "Soft_delete: true"`
}

func (User) SoftDelete() bool {
	return false
}

type UserServer struct {
	DB *gorm.DB
}
func (s *UserServer) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
    id := req.GetId()

    var user User
    res := s.DB.First(&user, id)
    if res.Error != nil {
        return &pb.UpdateUserResponse{Success: false}, res.Error
    }

    if req.DiamondsTotals != 0 {
        user.Diamonds_totals = strconv.FormatInt(req.GetDiamondsTotals(), 10)
    }
    if req.Fullname != "" {
        user.Fullname = req.GetFullname()
    }
    if req.Username != "" {
        user.Username = req.GetUsername()
    }
    if req.Email != "" {
        user.Email = req.GetEmail()
    }

    res = s.DB.Save(&user)
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
