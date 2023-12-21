package main

import (
	"automatix/app/gate/cmd/api/internal/config"
	"automatix/app/gate/cmd/api/internal/handler"
	"automatix/app/gate/cmd/api/internal/logic"
	"automatix/app/gate/cmd/api/internal/svc"
	"flag"
	"fmt"

	"automatix/common/net/zdecoder"
	"automatix/common/net/ziface"
	"automatix/common/net/znet"
	"automatix/common/net/zpack"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/gate.yaml", "the config file")

func OnConnectionAdd(conn ziface.IConnection) {
	fmt.Println("Connection is start")
	clientConn := logic.NewClientConn(conn)
	svc.ServiceContextObj.ClientConnManager.AddClientConn(clientConn)
}

func OnConnectionLost(conn ziface.IConnection) {

	fmt.Println("Connection is Lost")
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// s := svc.NewServer(ctx)
	handler.RegisterHandlers(ctx)
	// // Start Service
	// s.Serve()

	s := znet.NewServer()
	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	// Add LTV data format Decoder
	s.SetDecoder(zdecoder.NewLTV_Little_Decoder())
	// Add LTV data format Pack packet Encoder
	s.SetPacket(zpack.NewDataPackLtv())

	s.Serve()
}
