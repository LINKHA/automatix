package svc

import (
	"fmt"
	"net"
	"os"
)

func newTcpServer(ctx *ServiceContext) {
	// TCP Server
	address := fmt.Sprintf("%s:%d", ctx.Config.TcpHost, ctx.Config.TcpPort)

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println("Error resolving TCP address:", err)
		os.Exit(1)
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error listening for TCP:", err)
		os.Exit(1)
	}
	defer tcpListener.Close()

	fmt.Println("TCP server is running on", tcpAddr)

	go func() {
		for {
			conn, err := tcpListener.Accept()
			if err != nil {
				fmt.Println("Error accepting TCP connection:", err)
				continue
			}

			go handleTCPConnection(ctx, conn)
		}
	}()
	select {}
}

func newUdpServer(ctx *ServiceContext) {
	// UDP Server
	address := fmt.Sprintf("%s:%d", ctx.Config.UdpHost, ctx.Config.UdpPort)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	udpListener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening for UDP:", err)
		os.Exit(1)
	}
	defer udpListener.Close()

	fmt.Println("UDP server is running on", udpAddr)

	handleUDPConnection(udpListener)
	select {}
}

func NewGateServer(ctx *ServiceContext) {
	go newTcpServer(ctx)

	newUdpServer(ctx)

}
