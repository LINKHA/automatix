package svc

import (
	"automatix/app/rolemanager/cmd/rpc/pb"
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type ServerConfig struct {
	Name string
	IP   string
	Port int
}

type Server struct {
	SvcCtx   *ServiceContext
	Name     string
	IP       string
	Port     int
	exitChan chan struct{}

	ConnMgr *ConnManager
}

func NewServer(ctx *ServiceContext, config *ServerConfig) *Server {
	s := &Server{
		SvcCtx:   ctx,
		Name:     config.Name,
		IP:       config.IP,
		Port:     config.Port,
		exitChan: nil,

		ConnMgr: newConnManager(),
	}

	// //增加一个机制/每个server对应于每个服务都有一个client
	// method := reflect.ValueOf(ctx.RolemanagerRpc).MethodByName("CreateRoleStream")
	// // var client
	// // 检查方法是否存在
	// if method.IsValid() {
	// 	// 调用方法
	// 	method.Call(nil)
	// } else {
	// 	fmt.Println("Invalid method name:", methodName)
	// }

	client, _ := ctx.RolemanagerRpc.CreateRoleStream(context.Background())

	client.Send(&pb.CreateRoleReq{
		Id:        "1",
		Name:      "a",
		AccountId: 1,
	})

	rec, _ := client.Recv()
	fmt.Printf("xxx 1------------------   :   %v\n", rec)
	return s
}

func (s *Server) GetConnMgr() *ConnManager {
	return s.ConnMgr
}

func (s *Server) StartConn(conn Connection) {
	conn.Start()
}

func (s *Server) ListenTcpConn() {
	// TCP Server
	address := fmt.Sprintf("%s:%d", s.IP, s.Port)

	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println("Error resolving TCP address:", err)
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error listening for TCP:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("TCP server is running on", tcpAddr)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting TCP connection:", err)
				continue
			}
			connId := strconv.FormatInt(int64(s.SvcCtx.Snowflake.Generate()), 10)
			dealConn := newServerConn(s, conn, connId)

			go s.StartConn(dealConn)

		}
	}()

	select {
	case <-s.exitChan:
		err := listener.Close()
		if err != nil {
			fmt.Printf("listener close err: %v\n", err)
		}
	}
}

func (s *Server) Start() {
	s.exitChan = make(chan struct{})
	go s.ListenTcpConn()
}

func (s *Server) Serve() {
	s.Start()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	fmt.Printf("[SERVE] server , name %s, Serve Interrupt, signal = %v\n", s.Name, sig)
}
