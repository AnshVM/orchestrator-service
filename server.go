package main

import (
	"context"
	"errors"
	"log"
	"net"

	user "github.com/AnshVM/orchestrator-service/User"
	"google.golang.org/grpc"
)

type server struct {
	user.UnimplementedUserServiceServer
}

func (s *server) GetUserByName(ctx context.Context, name *user.Username) (*user.User, error) {
	return nil, errors.New("not implemented yet")
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &server{})
	log.Printf("Server up on port 9000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
