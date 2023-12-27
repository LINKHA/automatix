package iface

type IGrpcConnManager interface {
	Add(IGrpcConnection)
	Remove(IGrpcConnection)
	Get(uint64) (IGrpcConnection, error)
	Get2(string) (IGrpcConnection, error)
}
