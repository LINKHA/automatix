package svc

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type Server struct {
	SvcCtx   *ServiceContext
	Name     string
	IP       string
	Port     int
	exitChan chan struct{}

	ConnMgr *ConnManager
}

func NewServer(ctx *ServiceContext) *Server {

	s := &Server{
		SvcCtx:   ctx,
		Name:     "",
		IP:       ctx.Config.TcpHost,
		Port:     ctx.Config.TcpPort,
		exitChan: nil,

		ConnMgr: newConnManager(),
	}
	client, _ := ctx.RolemanagerRpc.CreateRoleStream(context.Background())

	grpc_conn := NewGrpcConnection(client, "")
	fmt.Printf("xxx 0------------------   :   %v\n", 1)
	go grpc_conn.Start()

	// err1 := grpc_conn.Send(&pb.CreateRoleReq{
	// 	Id:        "1",
	// 	Name:      "a",
	// 	AccountId: 1,
	// })
	// fmt.Printf("xxx 0------------------   :   %v\n", err1)
	// rec, err := grpc_conn.Recv(new(pb.CreateRoleResp))
	// fmt.Printf("xxx 1------------------   :   %v\n", rec)
	// fmt.Printf("xxx 2------------------   :   %v\n", err)
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
