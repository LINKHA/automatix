package logic

import (
	"automatix/common/net/ziface"
	"sync"
)

type ClientConn struct {
	PID  int32              // conn PID
	Conn ziface.IConnection // Current player's connection
}

var PIDGen int32 = 1
var IDLock sync.Mutex

// NewPlayer Create a player object
func NewClientConn(conn ziface.IConnection) *ClientConn {
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
