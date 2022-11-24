package logic

import (
	"context"
	"go-zero-template/service/app/rpc/appclient"

	"github.com/pkg/errors"

	"go-zero-template/service/app/api/internal/svc"
	"go-zero-template/service/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserReply, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.AppRpc.GetUser(l.ctx, &appclient.IdRequest{
		Id: req.Id,
	})

	// 返回错误
	if err != nil {
		return nil, errors.Wrapf(err, "GetUser req: %+v", req)
	}

	return &types.GetUserReply{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
