package iface

type IGrpcConnection interface {
	GetConnID() uint64
	GetConnIdStr() string
	Start()
	Stop()
	Send(interface{}) error
	SendToQueue(interface{}) error
	StartReader()
	StartWriter()
}
