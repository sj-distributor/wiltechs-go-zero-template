// Code generated by goctl. DO NOT EDIT!
// Source: app.proto

package server

import (
	"context"

	"go-zero-template/service/app/rpc/app"
	"go-zero-template/service/app/rpc/internal/logic"
	"go-zero-template/service/app/rpc/internal/svc"
)

type AppServer struct {
	svcCtx *svc.ServiceContext
	app.UnimplementedAppServer
}

func NewAppServer(svcCtx *svc.ServiceContext) *AppServer {
	return &AppServer{
		svcCtx: svcCtx,
	}
}

func (s *AppServer) GetUser(ctx context.Context, in *app.IdRequest) (*app.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}
