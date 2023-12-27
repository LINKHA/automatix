package homestayBussiness

import (
	"net/http"

	"github.com/LINKHA/automatix/app/travel/cmd/api/internal/logic/homestayBussiness"
	"github.com/LINKHA/automatix/app/travel/cmd/api/internal/svc"
	"github.com/LINKHA/automatix/app/travel/cmd/api/internal/types"
	"github.com/LINKHA/automatix/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomestayBussinessListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBussinessListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBussiness.NewHomestayBussinessListLogic(r.Context(), ctx)
		resp, err := l.HomestayBussinessList(req)
		result.HttpResult(r, w, resp, err)
	}
}
