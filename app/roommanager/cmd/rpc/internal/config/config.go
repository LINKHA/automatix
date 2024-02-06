package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Id int64
	zrpc.RpcServerConf
	KqConfServerQueue kq.KqConf
}
