package main

import (
	"google.golang.org/grpc"
	"head-first/grpc-training/adder"
	api "head-first/grpc-training/api/proto"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}

	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
