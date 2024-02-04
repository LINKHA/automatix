package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	mNet "github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	"github.com/LINKHA/automatix/app/gate/example/tcp/c_router"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	"google.golang.org/protobuf/proto"

	"github.com/LINKHA/automatix/common/log"
)

// Custom business logic of the client (客户端自定义业务)
func business(conn iface.IConnection) {
	data := &pb.RegisterRoleReq{
		AccountId:     "1",
		ServerId:      "1",
		TemplateValue: "{}",
	}
	msg, _ := proto.Marshal(data)

	for {
		// err := conn.SendMsg(1, []byte("Ping...[FromClient]"))
		err := conn.SendMsg(1, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
			break
		}

		time.Sleep(1 * time.Second)
	}
}

// Function to execute when the connection is created (创建连接的时候执行)
func DoClientConnectedBegin(conn iface.IConnection) {
	log.Debug("DoConnecionBegin is Called ... ")

	// Set two connection properties after the connection is created (设置两个链接属性，在连接创建之后)
	conn.SetProperty("Name", "刘丹冰Aceld")
	conn.SetProperty("Home", "https://yuque.com/aceld")

	go business(conn)
}

// Function to execute when the connection is lost (连接断开的时候执行)
func DoClientConnectedLost(conn iface.IConnection) {
	// Get the Name and Home properties of the connection before it is destroyed
	// (在连接销毁之前，查询conn的Name，Home属性)
	if name, err := conn.GetProperty("Name"); err == nil {
		log.Debug("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		log.Debug("Conn Property Home = ", home)
	}

	log.Debug("DoClientConnectedLost is Called ... ")
}

func main() {
	// Create a client handle using Zinx's Method (创建一个Client句柄，使用Zinx的方法)
	client := mNet.NewClient("127.0.0.1", 8999)

	// Set the business logic to execute when the connection is created or lost
	// (添加首次建立链接时的业务)
	client.SetOnConnStart(DoClientConnectedBegin)
	client.SetOnConnStop(DoClientConnectedLost)

	// Register routers for the messages received from the server
	// (注册收到服务器消息业务路由)
	client.AddRouter(2, &c_router.PingRouter{})
	client.AddRouter(3, &c_router.HelloRouter{})

	// Start the client
	client.Start()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Println("===exit===", sig)
	client.Stop()
	time.Sleep(time.Second * 2)
}
