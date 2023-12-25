package handler

import (
	"automatix/app/gate/cmd/api/internal/logic"
	"automatix/common/net/zconf"
	"automatix/common/net/ziface"
)

type MsgHandler struct {
	GrpcConnManager   *logic.GrpcConnManager
	ClientConnManager *logic.ClientConnManager

	grpcTaskQueue   []chan GrpcMessage
	clientTaskQueue []chan ClientMessage
}

func NewMsgHandler() *MsgHandler {
	grpcConnManager := logic.NewGrpcConnManager()
	clientConnManager := logic.NewClientConnManager()

	return &MsgHandler{
		GrpcConnManager:   grpcConnManager,
		ClientConnManager: clientConnManager,
		grpcTaskQueue:     make(chan GrpcMessage, 10),
		clientTaskQueue:   make(chan ClientMessage, 10),
	}
}

func (s *MsgHandler) HandleGrpc(msg GrpcMessage) {
	s.grpcTaskQueue <- msg
}

func (s *MsgHandler) HandleClient(msg ClientMessage) {
	s.clientTaskQueue <- msg
}

func (s *MsgHandler) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	// Continuously wait for messages in the queue
	// (不断地等待队列中的消息)
	for {
		select {
		// If there is a message, take out the Request from the queue and execute the bound business method
		// (有消息则取出队列的Request，并执行绑定的业务方法)
		case request := <-taskQueue:

			switch req := request.(type) {

			case ziface.IFuncRequest:
				// Internal function call request (内部函数调用request)

				mh.doFuncHandler(req, workerID)

			case ziface.IRequest: // Client message request

				if !zconf.GlobalObject.RouterSlicesMode {
					mh.doMsgHandler(req, workerID)
				} else if zconf.GlobalObject.RouterSlicesMode {
					mh.doMsgHandlerSlices(req, workerID)
				}
			}
		}
	}
}
