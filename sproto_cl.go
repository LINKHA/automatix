package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/linkha/automatix/apigrpc" // 导入您的生成的 gRPC 代码包
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	// 建立 gRPC 连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// 创建 GreetService 客户端
	client := apigrpc.NewGreetServiceClient(conn)

	// 准备请求消息
	request := &apigrpc.HelloRequest{
		Name: "YourName",
	}

	// 调用 gRPC 服务中的 SayHello 方法
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, request)
	if err != nil {
		// 处理 gRPC 调用错误
		if status.Code(err) == codes.DeadlineExceeded {
			fmt.Println("Request timeout")
		} else {
			log.Fatalf("Failed to call SayHello: %v", err)
		}
		return
	}

	// 处理响应
	fmt.Printf("Response: %s\n", response.Message)
}
