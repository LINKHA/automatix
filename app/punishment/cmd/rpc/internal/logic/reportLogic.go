package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/punishment/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/punishment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportLogic {
	return &ReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReportLogic) Report(stream pb.Punishment_ReportServer) error {
	// todo: add your logic here and delete this line

	return nil
}
