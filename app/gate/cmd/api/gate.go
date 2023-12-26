package main

import (
	"automatix/app/gate/cmd/api/internal/config"
	"automatix/app/gate/cmd/api/internal/handler"
	"automatix/app/gate/cmd/api/internal/logic"
	"flag"
	"fmt"

	"automatix/app/gate/cmd/api/internal/net/decoder"
	"automatix/app/gate/cmd/api/internal/net/iface"
	"automatix/app/gate/cmd/api/internal/net/net"
	"automatix/app/gate/cmd/api/internal/net/pack"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/gate.yaml", "the config file")

func OnConnectionAdd(conn iface.IConnection) {
	clientConn := logic.NewClientConn(conn)
	handler.ServiceContextObj.MsgHandler.ClientConnManager.AddConn(clientConn)

	conn.SetProperty("pID", clientConn.PID)
	fmt.Println("Connection is start")
}

func OnConnectionLost(conn iface.IConnection) {
	pID, _ := conn.GetProperty("pId")
	var connId int32

	if pID != nil {
		connId = pID.(int32)
		handler.ServiceContextObj.MsgHandler.ClientConnManager.RemoveConnByPID(connId)
	}
	fmt.Println("Connection is Lost. connId: ", connId)
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := handler.NewServiceContext(c)

	handler.RegisterHandlers(ctx)

	s := net.NewServer()
	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	// Add LTV data format Decoder
	s.SetDecoder(decoder.NewLTV_Little_Decoder())
	// Add LTV data format Pack packet Encoder
	s.SetPacket(pack.NewDataPackLtv())

	s.Serve()
}
