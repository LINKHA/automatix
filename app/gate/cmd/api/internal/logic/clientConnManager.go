package logic

import "sync"

type ClientConnManager struct {
	ClientConns map[int32]*ClientConn
	pLock       sync.RWMutex
}

var ClientConnMgrObj *ClientConnManager

func NewClientConnManager() *ClientConnManager {
	return &ClientConnManager{
		ClientConns: make(map[int32]*ClientConn),
	}
}

func (ccm *ClientConnManager) AddClientConn(clientConn *ClientConn) {
	// Add the player to the world manager
	// 将player添加到 世界管理器中
	ccm.pLock.Lock()
	ccm.ClientConns[clientConn.PID] = clientConn
	ccm.pLock.Unlock()
}

func (ccm *ClientConnManager) RemoveClientConnByPID(pID int32) {
	ccm.pLock.Lock()
	delete(ccm.ClientConns, pID)
	ccm.pLock.Unlock()
}
