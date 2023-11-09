package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Id int64
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	Cache             cache.CacheConf
	KqConfServerQueue kq.KqConf
}
