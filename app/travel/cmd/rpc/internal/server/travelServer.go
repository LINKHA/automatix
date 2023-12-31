// Code generated by goctl. DO NOT EDIT!
// Source: travel.proto

package server

import (
	"context"

	"github.com/LINKHA/automatix/app/travel/cmd/rpc/internal/logic"
	"github.com/LINKHA/automatix/app/travel/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/travel/cmd/rpc/pb"
)

type TravelServer struct {
	svcCtx *svc.ServiceContext
}

func NewTravelServer(svcCtx *svc.ServiceContext) *TravelServer {
	return &TravelServer{
		svcCtx: svcCtx,
	}
}

// 民宿详情
func (s *TravelServer) HomestayDetail(ctx context.Context, in *pb.HomestayDetailReq) (*pb.HomestayDetailResp, error) {
	l := logic.NewHomestayDetailLogic(ctx, s.svcCtx)
	return l.HomestayDetail(in)
}
