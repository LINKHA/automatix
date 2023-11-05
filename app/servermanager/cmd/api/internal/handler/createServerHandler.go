package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/servermanager/cmd/api/internal/logic"
	"looklook/app/servermanager/cmd/api/internal/svc"
	"looklook/app/servermanager/cmd/api/internal/types"
)

func createServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateServerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateServerLogic(r.Context(), svcCtx)
		resp, err := l.CreateServer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
