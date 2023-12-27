package user

import (
	"net/http"

	"github.com/LINKHA/automatix/app/usercenter/cmd/api/internal/logic/user"
	"github.com/LINKHA/automatix/app/usercenter/cmd/api/internal/svc"
	"github.com/LINKHA/automatix/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FastLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FastLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewFastLoginLogic(r.Context(), svcCtx)
		resp, err := l.FastLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
