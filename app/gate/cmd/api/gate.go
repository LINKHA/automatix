package main

import (
	"automatix/app/gate/cmd/api/internal/config"
	"automatix/app/gate/cmd/api/internal/handler"
	"automatix/app/gate/cmd/api/internal/logic"
	"flag"
	"fmt"

	"automatix/app/gate/cmd/api/internal/net/zdecoder"
	"automatix/app/gate/cmd/api/internal/net/ziface"
	"automatix/app/gate/cmd/api/internal/net/znet"
	"automatix/app/gate/cmd/api/internal/net/zpack"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/gate.yaml", "the config file")

func OnConnectionAdd(conn ziface.IConnection) {
	clientConn := logic.NewClientConn(conn)
	handler.ServiceContextObj.MsgHandler.ClientConnManager.AddConn(clientConn)

	conn.SetProperty("pID", clientConn.PID)
	fmt.Println("Connection is start")
}

func OnConnectionLost(conn ziface.IConnection) {
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

	s := znet.NewServer()
	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	// Add LTV data format Decoder
	s.SetDecoder(zdecoder.NewLTV_Little_Decoder())
	// Add LTV data format Pack packet Encoder
	s.SetPacket(zpack.NewDataPackLtv())

	s.Serve()
}
