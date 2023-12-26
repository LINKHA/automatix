package logic

import (
	"automatix/app/gate/cmd/api/internal/net/iface"
	"sync"
)

type ClientConn struct {
	PID  int32
	Conn iface.IConnection
}

var PIDGen int32 = 1
var IDLock sync.Mutex

func NewClientConn(conn iface.IConnection) *ClientConn {
	IDLock.Lock()
	ID := PIDGen
	PIDGen++
	IDLock.Unlock()

	p := &ClientConn{
		PID:  ID,
		Conn: conn,
	}

	return p
}

func (p *ClientConn) SyncPID() {

}
