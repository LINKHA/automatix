package logic

import (
	"sync"
)

type GrpcConn struct {
	PID  int32
	Conn IGrpcConnection
}

var GrpcPIDGen int32 = 1
var GrpcIDLock sync.Mutex

func NewGrpcConn(conn IGrpcConnection) *GrpcConn {
	IDLock.Lock()
	ID := PIDGen
	PIDGen++
	IDLock.Unlock()

	p := &GrpcConn{
		PID:  ID,
		Conn: conn,
	}

	return p
}

func (p *GrpcConn) SyncPID() {

}
