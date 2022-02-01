package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	orchestrator "github.com/AnshVM/orchestrator-service/logic"

	user "github.com/AnshVM/orchestrator-service/User"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

type server struct {
	user.UnimplementedUserServiceServer
}

func (s *server) GetUserByName(ctx context.Context, name *user.Username) (*user.User, error) {
	log.Printf("Recieved req on port 9000 GetUserByName service")
	newUser, err := orchestrator.Forward_request("localhost:9001", "GetUser", name)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *server) GetUser(ctx context.Context, name *user.Username) (*user.User, error) {
	log.Printf("Recieved req on port 9001 GetUser service")
	newUser, err := orchestrator.Forward_request("localhost:10000", "GetMockUserData", name)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", *port, err)
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &server{})
	log.Printf("Server up on port %d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
