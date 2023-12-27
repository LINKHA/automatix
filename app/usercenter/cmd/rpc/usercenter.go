package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/service"
	"google.golang.org/grpc/reflection"

	"github.com/LINKHA/automatix/app/usercenter/cmd/rpc/internal/config"
	"github.com/LINKHA/automatix/app/usercenter/cmd/rpc/internal/server"
	"github.com/LINKHA/automatix/app/usercenter/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/usercenter/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/interceptor/rpcserver"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUsercenterServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, srv)
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
