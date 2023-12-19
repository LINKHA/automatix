package svc

type IConnection interface {
	Start()
	Stop()
	GetConnID() string
}
