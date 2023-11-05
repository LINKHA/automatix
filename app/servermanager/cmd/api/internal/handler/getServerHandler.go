package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/servermanager/cmd/api/internal/logic"
	"looklook/app/servermanager/cmd/api/internal/svc"
	"looklook/app/servermanager/cmd/api/internal/types"
)

func getServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetServerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetServerLogic(r.Context(), svcCtx)
		resp, err := l.GetServer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
