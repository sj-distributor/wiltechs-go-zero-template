package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"go-zero-template/common/swagger"
	"go-zero-template/service/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Swag   http.HandlerFunc
}

const swagFilePath = "./doc/app.json"

var SwagByte json.RawMessage

func init() {
	swagFile, err := os.Open(swagFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer swagFile.Close()
	SwagByte, err = ioutil.ReadAll(swagFile)
	if err != nil {
		fmt.Println(err)
	}
}

func NewSwagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwagLogic {
	return &SwagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		Swag:   swagger.Doc("/swagger", SwagByte),
	}
}
