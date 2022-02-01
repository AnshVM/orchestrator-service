//newUser = forward("9001","GetUser",name)
//newUser = forward("10000","GetMockUserData",name)

package logic

import (
	"context"
	"errors"
	"log"
	"time"

	user "github.com/AnshVM/orchestrator-service/User"
	dummy "github.com/AnshVM/orchestrator-service/datamock/Dummy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Forward_request(port string, rpcService string, username *user.Username) (*user.User, error) {
	if rpcService == "GetUser" {
		conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}
		defer conn.Close()
		client := user.NewUserServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		user, err := client.GetUser(ctx, username)
		if err != nil {
			log.Fatalf("RPC service GetUser failed")
		}
		return user, nil
	}

	if rpcService == "GetMockUserData" {
		conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}
		defer conn.Close()
		client := dummy.NewDummyDataServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		dummyUser, err := client.GetMockUserData(ctx, &dummy.DummyUsername{Name: username.GetName()})
		if err != nil {
			log.Fatalf("RPC service GetUser failed")
		}
		user := user.User{
			Name:  dummyUser.GetName(),
			Class: dummyUser.GetClass(),
			Roll:  dummyUser.GetRoll(),
		}
		return &user, nil
	}

	return nil, errors.New("invalid input")
}
