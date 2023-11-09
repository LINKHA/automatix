package user

import (
	"net/http"

	"automatix/app/usercenter/cmd/api/internal/logic/user"
	"automatix/app/usercenter/cmd/api/internal/svc"
	"automatix/app/usercenter/cmd/api/internal/types"
	"automatix/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewDetailLogic(r.Context(), ctx)
		resp, err := l.Detail(req)
		result.HttpResult(r, w, resp, err)
	}
}
