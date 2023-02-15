package handler

import (
	"context"
	"net/http"

	"go-zero-template/service/app/api/internal/logic"
	"go-zero-template/service/app/api/internal/svc"
)

func swagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	l := logic.NewSwagLogic(context.TODO(), svcCtx)
	return l.Swag
}
