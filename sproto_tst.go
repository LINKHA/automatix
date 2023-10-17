package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/linkha/automatix/apigrpc" // Import the generated protobuf package

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + in.Name}, nil
}

// // 实现 mustEmbedUnimplementedGreetServiceServer 方法
//
//	func (s *server) mustEmbedUnimplementedGreetServiceServer() {
//		// 实现具体逻辑
//	}
//
// grpcurl -plaintext -d '{"message": "Hello from gRPC!"}' localhost:50051 example.Greeter/SayHello
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	// ss := &server{}

	pb.RegisterGreetServiceServer(grpcServer, pb.UnimplementedGreetServiceServer{})

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
