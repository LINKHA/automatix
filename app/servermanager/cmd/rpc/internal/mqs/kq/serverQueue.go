package kq

import (
	"automatix/app/servermanager/cmd/rpc/internal/svc"
	"automatix/common/kqueue"
	"automatix/common/servercode"
	"context"
	"encoding/json"

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
	//登录
	servercode.GenServerCode(l.svcCtx.Redis, message.UserId, message.ServerId)
	return nil
}
