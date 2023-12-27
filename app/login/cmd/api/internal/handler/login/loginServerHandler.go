package login

import (
	"net/http"

	"github.com/LINKHA/automatix/app/login/cmd/api/internal/logic/login"
	"github.com/LINKHA/automatix/app/login/cmd/api/internal/svc"
	"github.com/LINKHA/automatix/app/login/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginServerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewLoginServerLogic(r.Context(), svcCtx)
		resp, err := l.LoginServer(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
