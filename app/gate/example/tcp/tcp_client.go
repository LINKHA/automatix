package tcp

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:11111")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("Hello from client %d", i)
		// 向服务器发送数据
		conn.Write([]byte(message))

		// 读取服务器的响应
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		receivedData := buffer[:n]
		fmt.Printf("Received from server: %s\n", receivedData)
	}
}
