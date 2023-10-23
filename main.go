go
package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "google.golang.org/grpc"
    "your_package_path/upper" // 导入生成的 protobuf 包
)

func main() {
    // 创建 Gin 引擎
    router := gin.Default()

    // 设置 HTTP 路由，将不同的路径映射到不同的处理函数
    router.POST("/", handleHTTPForward)
    // 添加更多的 HTTP 路由规则，如果需要的话

    // 同一个端口上同时启动 HTTP 和 gRPC 服务器
    go startHTTPServer(router)

    // 保持程序运行
    select {}
}

func startHTTPServer(router *gin.Engine) {
    httpAddr := ":8080" // HTTP 和 gRPC 服务器共享的端口
    log.Printf("Server is listening on %s\n", httpAddr)
    if err := router.Run(httpAddr); err != nil {
        log.Fatalf("HTTP server failed: %v", err)
    }
}

func handleHTTPForward(c *gin.Context) {
    // 在这里实现 HTTP 请求的处理和转发逻辑
    // 你可以将 HTTP 请求转发到 gRPC 服务，或者其他目标服务
    // 示例：HTTP 请求处理逻辑

    text := c.PostForm("text") // 假设从 HTTP 请求中获取文本数据
    if text == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Text is missing"})
        return
    }

    // 连接到远程的 gRPC 服务
    conn, err := grpc.Dial("grpc-server-address:50051", grpc.WithInsecure())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
        return
    }
    defer conn.Close()

    client := upper.NewUpperServiceClient(conn)
    // 调用 gRPC 服务
    response, err := client.ToUpper(context.Background(), &upper.UpperRequest{Text: text})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call gRPC service"})
        return
    }

    // 将 gRPC 服务的响应返回给 HTTP 客户端
    c.JSON(http.StatusOK, gin.H{"upper_text": response.UpperText})
}