package adapter

import "context"

type GRPCServer struct{}

func (s *GRPCServer) Health(context.Context, *struct{}) (*struct{ Status string }, error) {
	return &struct{ Status string }{Status: "ok"}, nil
}
