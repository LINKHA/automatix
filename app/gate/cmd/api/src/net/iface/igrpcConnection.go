package iface

type IGrpcConnection interface {
	GetConnId() uint64
	GetConnIdStr() string
	Start()
	Stop()
	Send(interface{}) error
	SendToReqQueue(IRequest) error
	StartReader()
	StartWriter()
}
