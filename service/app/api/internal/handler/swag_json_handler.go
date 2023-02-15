package handler

import (
	"net/http"

	"go-zero-template/service/app/api/internal/logic"
	"go-zero-template/service/app/api/internal/svc"
)

func swagJsonHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write(logic.SwagByte)
	}
}
