package adder

import (
	"context"
	"head-first/grpc training/api/proto"
)

type GRPCServer struct {
}

func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {
	//TODO implement me

	panic("implement me")
}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
