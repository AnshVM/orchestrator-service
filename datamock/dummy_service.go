package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	dummy "github.com/AnshVM/orchestrator-service/datamock/Dummy"
	"google.golang.org/grpc"
)

type server struct {
	dummy.UnimplementedDummyDataServiceServer
}

func (s *server) GetMockUserData(ctx context.Context, name *dummy.DummyUsername) (*dummy.DummyUser, error) {
	log.Printf("Recieved req on port 10000 GetMockUserData service")

	if len(name.GetName()) < 6 {
		return nil, errors.New("username is less than 6 characters")
	}
	user := dummy.DummyUser{
		Name:  name.GetName(),
		Class: strconv.Itoa(len(name.GetName())),
		Roll:  int64(len(name.GetName()) * 10),
	}
	return &user, nil
}

func main() {
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("Failed to listen on port 10000: %v", err)
	}
	s := grpc.NewServer()
	dummy.RegisterDummyDataServiceServer(s, &server{})
	log.Printf("Server up on port 10000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
