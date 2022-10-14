package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-template/service/app/api/internal/config"
	"go-zero-template/service/app/rpc/appclient"
)

type ServiceContext struct {
	Config config.Config
	AppRpc appclient.App
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AppRpc: appclient.NewApp(zrpc.MustNewClient(c.AppRpc)),
	}
}
