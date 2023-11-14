package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	api "head-first/grpc-training/api/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewAdderClient(conn)
	res, err := client.Add(context.TODO(), &api.AddRequest{
		X: 2,
		Y: 5,
	})
	if err != nil {
		return
	}
	log.Println(res.GetResult())
}
