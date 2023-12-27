package handler

import (
	"net/http"

	"github.com/LINKHA/automatix/app/servermanager/cmd/api/internal/logic"
	"github.com/LINKHA/automatix/app/servermanager/cmd/api/internal/svc"
	"github.com/LINKHA/automatix/app/servermanager/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getServerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetServerListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetServerListLogic(r.Context(), svcCtx)
		resp, err := l.GetServerList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
