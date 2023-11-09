package handler

import (
	"net/http"

	"automatix/app/servermanager/cmd/api/internal/logic"
	"automatix/app/servermanager/cmd/api/internal/svc"
	"automatix/app/servermanager/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func setServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetServerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSetServerLogic(r.Context(), svcCtx)
		resp, err := l.SetServer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
