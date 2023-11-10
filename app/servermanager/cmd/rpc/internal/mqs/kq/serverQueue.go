package kq

import (
	"automatix/app/servermanager/cmd/rpc/internal/svc"
	"automatix/common/kqueue"
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServerQueueMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewServerQueueMq(ctx context.Context, svcCtx *svc.ServiceContext) *ServerQueueMq {
	return &ServerQueueMq{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ServerQueueMq) Consume(_, val string) error {
	fmt.Print("qwe 3---------    :   ")
	fmt.Println("consume login msg")
	fmt.Println(val)

	var message kqueue.LoginServerMessage

	if err := json.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}

	if err := l.execService(message); err != nil {
		logx.WithContext(l.ctx).Error("PaymentUpdateStatusMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *ServerQueueMq) execService(message kqueue.LoginServerMessage) error {
	// serverId = message.ServerId
	// l.svcCtx.Redis.Incr(message.ServerId + )
	return nil
}
