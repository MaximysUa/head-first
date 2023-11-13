package main

import (
	"google.golang.org/grpc"
	"head-first/grpc training/adder"
	api "head-first/grpc training/api/proto"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)
}
