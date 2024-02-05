package iface

type IGrpcConnection interface {
	GetConnId() uint64
	GetConnIdStr() string
	Start()
	Stop()
	Send(interface{}) error
	SendToReqQueue(uint64, []byte) error
	StartReader()
	StartWriter()
}
