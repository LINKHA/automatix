package iface

type IGrpcConnection interface {
	GetConnId() uint64
	GetConnIdStr() string
	Start()
	Stop()
	Send(interface{}) error
	SendToQueue([]byte) error
	StartReader()
	StartWriter()
}
