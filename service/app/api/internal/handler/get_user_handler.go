package handler

import (
	"go-zero-template/common/result"
	"net/http"

	"go-zero-template/service/app/api/internal/logic"
	"go-zero-template/service/app/api/internal/svc"
	"go-zero-template/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		result.HttpResult(r, w, resp, err)
	}
}
