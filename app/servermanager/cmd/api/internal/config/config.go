package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type KqConfig struct {
	Brokers []string
	Topic   string
}

type Config struct {
	rest.RestConf
	ServerManagerRpcConf zrpc.RpcClientConf
	Id                   int64
	KqConfServerQueue    KqConfig
	JwtAuth              struct {
		AccessSecret string
	}
}
