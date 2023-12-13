package svc

import (
	"fmt"
	"net"
	"os"
	"strconv"
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

	return s
}

func (s *Server) GetConnMgr() *ConnManager {
	return s.ConnMgr
}

func (s *Server) StartConn(conn Connection) {
	// HeartBeat check
	if s.hc != nil {
		// Clone a heart-beat checker from the server side
		heartBeatChecker := s.hc.Clone()

		// Bind current connection
		heartBeatChecker.BindConn(conn)
	}

	// Start processing business for the current connection
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
			fmt.Println("listener close err: %v", err)
		}
	}
}
