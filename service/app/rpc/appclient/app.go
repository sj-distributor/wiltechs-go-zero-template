// Code generated by goctl. DO NOT EDIT!
// Source: app.proto

package appclient

import (
	"context"

	"go-zero-template/service/app/rpc/app"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest    = app.IdRequest
	UserResponse = app.UserResponse

	App interface {
		GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	}

	defaultApp struct {
		cli zrpc.Client
	}
)

func NewApp(cli zrpc.Client) App {
	return &defaultApp{
		cli: cli,
	}
}

func (m *defaultApp) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	client := app.NewAppClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}
