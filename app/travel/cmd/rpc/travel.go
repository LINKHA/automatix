package main

import (
	"flag"
	"fmt"

	"github.com/LINKHA/automatix/app/travel/cmd/rpc/internal/config"
	"github.com/LINKHA/automatix/app/travel/cmd/rpc/internal/server"
	"github.com/LINKHA/automatix/app/travel/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/travel/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/travel.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewTravelServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterTravelServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
