package svc

import (
	"github.com/LINKHA/automatix/app/order/cmd/mq/internal/config"
	"github.com/LINKHA/automatix/app/order/cmd/rpc/order"
	"github.com/LINKHA/automatix/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc      order.Order
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
