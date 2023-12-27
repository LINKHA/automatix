package config

import "github.com/zeromicro/go-zero/zrpc"

type GateConfig struct {
	TcpHost string `json:",default=0.0.0.0"`
	TcpPort int
	UdpHost string `json:",default=0.0.0.0"`
	UdpPort int
}

type Config struct {
	// rest.RestConf
	GateConfig
	RoleManagerRpcConf zrpc.RpcClientConf
	RoomManagerRpcConf zrpc.RpcClientConf
	Id                 int64
}
