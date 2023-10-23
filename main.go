package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux" // 使用Gorilla Mux来处理HTTP路由
	"google.golang.org/grpc" // 导入gRPC库
	// 导入其他必要的库
)

func main() {
	// 创建HTTP路由器
	router := mux.NewRouter()

	// 设置HTTP路由，将不同的路径映射到不同的处理函数
	router.HandleFunc("/httpEndpoint", handleHTTPForward).Methods("POST")
	// 添加更多的HTTP路由规则，如果需要的话

	// 启动HTTP服务器
	go startHTTPServer(router)

	// 连接到gRPC服务器
	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcConn.Close()

	// 启动gRPC服务
	go startGRPCServer(grpcConn)

	// 保持程序运行
	select {}
}

func startHTTPServer(router *mux.Router) {
	httpAddr := ":8080" // HTTP服务器的地址
	log.Printf("HTTP server is listening on %s\n", httpAddr)
	if err := http.ListenAndServe(httpAddr, router); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}

func startGRPCServer(grpcConn *grpc.ClientConn) {
	// 在这里实现gRPC服务的逻辑
	// 你可以使用protobuf生成的代码来处理gRPC请求和响应
	// 示例：实现一个gRPC函数
}

func handleHTTPForward(w http.ResponseWriter, r *http.Request) {
	// 在这里实现HTTP请求的处理和转发逻辑
	// 你可以将HTTP请求转发到gRPC服务，或者其他目标服务
	// 示例：HTTP请求处理逻辑
}
