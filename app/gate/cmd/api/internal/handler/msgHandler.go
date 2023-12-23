package handler

import (
	"automatix/app/gate/cmd/api/internal/logic"
	"automatix/common/net/zlog"
	"fmt"
)

type MsgHandler struct {
	Apis            map[uint32]logic.IGrpcConnection
	GrpcTaskQueue   chan GrpcMessage
	ClientTaskQueue chan ClientMessage
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:            make(map[uint32]logic.IGrpcConnection),
		GrpcTaskQueue:   make(chan GrpcMessage, 10),
		ClientTaskQueue: make(chan ClientMessage, 10),
	}
}

func (mh *MsgHandler) AddConn(msgID uint32, grpcConn logic.IGrpcConnection) {
	// 1. Check whether the current API processing method bound to the msgID already exists
	// (判断当前msg绑定的API处理方法是否已经存在)
	if _, ok := mh.Apis[msgID]; ok {
		msgErr := fmt.Sprintf("repeated api , msgID = %+v\n", msgID)
		panic(msgErr)
	}
	// 2. Add the binding relationship between msg and API
	// (添加msg与api的绑定关系)
	mh.Apis[msgID] = grpcConn
	zlog.Ins().InfoF("Add Router msgID = %d", msgID)
}

func (mh *MsgHandler) Handle() {

}
