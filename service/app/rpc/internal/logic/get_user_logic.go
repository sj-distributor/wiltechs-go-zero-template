package logic

import (
	"context"
	"go-zero-template/common/xerr"

	"github.com/pkg/errors"

	"go-zero-template/service/app/rpc/app"
	"go-zero-template/service/app/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *app.IdRequest) (*app.UserResponse, error) {
	// todo: add your logic here and delete this line

	// 返回错误
	if in.Id != "1" {
		return nil, errors.Wrapf(xerr.NewErrMsg("错误id"), "GetUser error")
	}

	return &app.UserResponse{
		Id:   "1",
		Name: "test",
	}, nil
}
