package main

import (
	"context"
	"log"
	"time"

	user "github.com/AnshVM/orchestrator-service/User"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	client := user.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetUserByName(ctx, &user.Username{Name: "Ansh"})
	if err != nil {
		log.Fatalf("Recieved error: %v", err)
	}
	log.Printf("%v", r.GetName())
}
