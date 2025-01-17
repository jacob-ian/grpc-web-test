package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"jacobmatthe.ws/grpc-web-test/protos/dashboard"
)

type service struct {
	dashboard.UnimplementedDashboardServer
}

func (s *service) GetGreeting(ctx context.Context, r *dashboard.GetGreetingRequest) (*dashboard.GetGreetingResponse, error) {
	return &dashboard.GetGreetingResponse{
		Greeting: &dashboard.Greeting{
			Id:      "37bed5cb-bc49-41d3-95cf-182469751da6",
			Message: "Hello there!",
		},
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	service := &service{}

	s := grpc.NewServer()
	dashboard.RegisterDashboardServer(s, service)

	println("listening")

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
