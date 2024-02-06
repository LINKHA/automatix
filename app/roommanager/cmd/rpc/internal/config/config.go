package config

import (
	"github.com/LINKHA/automatix/common/flowlimit"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Id int64
	zrpc.RpcServerConf
	Cache             cache.CacheConf
	KqConfServerQueue kq.KqConf
	SlidingWindow     flowlimit.SlidingWindowConf
}
