package svc

import (
	"context"
	"fmt"
	"net"
)

func handleTCPConnection(ctx *ServiceContext, conn net.Conn) {
	defer conn.Close()
	ctx.RolemanagerRpc.CreateRoleStream(context.Background())

	// roleStream, err := ctx.RolemanagerRpc.CreateRoleStream(context.Background())
	// fmt.Print("xx 1------------------   :   ")
	// fmt.Println(err)

	// roleStream.Send(&pb.CreateRoleReq{
	// 	Id:        "1",
	// 	Name:      "2",
	// 	AccountId: 123,
	// })

	// roleStream.Send(&pb.CreateRoleReq{
	// 	Id:        "1x",
	// 	Name:      "2x",
	// 	AccountId: 123,
	// })

	// ss, _ := roleStream.Recv()
	// ss2, _ := roleStream.Recv()
	// fmt.Print("xx 2------------------   :   ")
	// fmt.Println(ss)
	// fmt.Println(ss2)

	// buffer := make([]byte, 1024)
	// for {
	// 	n, err := conn.Read(buffer)
	// 	if err != nil {
	// 		fmt.Println("Error reading:", err)
	// 		return
	// 	}

	// 	fmt.Printf("Received TCP message: %s\n", string(buffer[:n]))

	// 	_, err = conn.Write(buffer[:n])
	// 	if err != nil {
	// 		fmt.Println("Error writing:", err)
	// 		return
	// 	}
	// }
}

func handleUDPConnection(conn *net.UDPConn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		fmt.Printf("Received UDP message from %s: %s\n", addr.String(), string(buffer[:n]))

		_, err = conn.WriteToUDP(buffer[:n], addr)
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}
