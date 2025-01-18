package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"jacobmatthe.ws/grpc-web-test/protos/dashboard"
)

func main() {
	wg := &sync.WaitGroup{}
	println("started")
	wg.Add(2)
	go runGrpc(":9090", wg)
	go runHttp(":9091", wg)
	wg.Wait()
}

func runHttp(addr string, wg *sync.WaitGroup) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("howdy"))
	})

	server := &http.Server{
		Addr: addr,
	}

	if err := server.ListenAndServe(); err != nil {
		wg.Done()
		panic(err)
	}
}

func runGrpc(addr string, wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		wg.Done()
		panic(err)
	}

	service := &service{}

	s := grpc.NewServer()
	dashboard.RegisterDashboardServer(s, service)
	reflection.Register(s)

	println("grpc listening")

	if err := s.Serve(lis); err != nil {
		wg.Done()
		panic(err)
	}
}

type service struct {
	dashboard.UnimplementedDashboardServer
}

func (s *service) GetGreeting(ctx context.Context, r *dashboard.GetGreetingRequest) (*dashboard.GetGreetingResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Printf("%#v\n", md)
	return &dashboard.GetGreetingResponse{
		Greeting: &dashboard.Greeting{
			Id:      "37bed5cb-bc49-41d3-95cf-182469751da6",
			Message: "Hello there!",
		},
	}, nil

}
