package login

import (
	"net/http"

	"automatix/app/login/cmd/api/internal/logic/login"
	"automatix/app/login/cmd/api/internal/svc"
	"automatix/app/login/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetServerCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetServerCodeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewGetServerCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetServerCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
