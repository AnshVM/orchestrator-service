package main

import (
	"context"
	"log"
	"time"

	dummy "github.com/AnshVM/orchestrator-service/datamock/Dummy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	client := dummy.NewDummyDataServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.GetMockUserData(ctx, &dummy.Username{Name: "Ansh"})
	if err != nil {
		log.Fatalf("Recieved error: %v", err)
	}
	log.Printf("%v", r)
}
