package main

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	// 导入其他必要的库
)

func main() {
	// 创建Gin引擎
	router := gin.Default()

	// 设置HTTP路由，将不同的路径映射到不同的处理函数
	router.POST("/httpEndpoint", handleHTTPForward)
	// 添加更多的HTTP路由规则，如果需要的话

	// 同一个端口上同时启动HTTP和gRPC服务器
	go startHTTPServer(router)
	go startGRPCServer()

	// 保持程序运行
	select {}
}

func startHTTPServer(router *gin.Engine) {
	httpAddr := ":8080" // HTTP和gRPC服务器共享的端口
	log.Printf("Server is listening on %s\n", httpAddr)
	if err := router.Run(httpAddr); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}

func startGRPCServer() {
	grpcAddr := ":8080" // HTTP和gRPC服务器共享的端口
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", grpcAddr, err)
	}

	grpcServer := grpc.NewServer()
	// 在这里注册gRPC服务

	log.Printf("gRPC server is listening on %s\n", grpcAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}

func handleHTTPForward(c *gin.Context) {
	// 在这里实现HTTP请求的处理和转发逻辑
	// 你可以将HTTP请求转发到gRPC服务，或者其他目标服务
	// 示例：HTTP请求处理逻辑
}
