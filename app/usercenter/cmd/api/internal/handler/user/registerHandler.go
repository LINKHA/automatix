package user

import (
	"net/http"

	"automatix/app/usercenter/cmd/api/internal/logic/user"
	"automatix/app/usercenter/cmd/api/internal/svc"
	"automatix/app/usercenter/cmd/api/internal/types"
	"automatix/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		result.HttpResult(r, w, resp, err)
	}
}
