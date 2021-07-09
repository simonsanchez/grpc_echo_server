package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/simonsanchez/grpc_echo_proto_two/v2/generated"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEchoServiceServer(grpcServer, &server{})
	fmt.Println("server listening...")
	grpcServer.Serve(lis)
}

type server struct {
	pb.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Printf("message recieved by server: %s\n", in.GetMessage())
	return &pb.EchoResponse{Message: in.GetMessage()}, nil
}

func (s *server) Foo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Println("foo -- " + time.Now().String())
	return &pb.EchoResponse{Message: in.GetMessage()}, nil
}

func (s *server) Bar(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Println("bar -- " + time.Now().String())
	return &pb.EchoResponse{Message: in.GetMessage()}, nil
}

func (s *server) Baz(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Println("baz -- " + time.Now().String())
	return &pb.EchoResponse{Message: in.GetMessage()}, nil
}
