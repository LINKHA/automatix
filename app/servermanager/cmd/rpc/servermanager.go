package main

import (
	"flag"
	"fmt"

	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/config"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/mqs/listen"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/server"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/servermanager.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	rpcServer := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterServermanagerServer(grpcServer, server.NewServermanagerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer rpcServer.Stop()

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	go rpcServer.Start()
	go serviceGroup.Start()

	select {}
}
