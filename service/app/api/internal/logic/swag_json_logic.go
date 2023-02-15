package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-template/service/app/api/internal/svc"
)

type SwagJsonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwagJsonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwagJsonLogic {
	return &SwagJsonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwagJsonLogic) SwagJson() error {
	// todo: add your logic here and delete this line

	return nil
}
