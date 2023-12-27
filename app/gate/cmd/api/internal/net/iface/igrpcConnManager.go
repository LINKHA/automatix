package iface

type IGrpcConnManager interface {
	Add(IGrpcConnection)
	Remove(string)
	Get(string) IGrpcConnection
}
