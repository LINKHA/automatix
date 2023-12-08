package udp

import (
	"fmt"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:11112")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP server:", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 5; i++ {
		message := fmt.Sprintf("Hello from UDP client %d", i)
		// 向服务器发送数据
		conn.Write([]byte(message))

		// 读取服务器的响应
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP server:", err)
			return
		}

		receivedData := buffer[:n]
		fmt.Printf("Received from UDP server: %s\n", receivedData)
	}
}
