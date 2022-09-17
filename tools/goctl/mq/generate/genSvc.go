package generate

import (
	"github.com/xiaoshouchen/go-zero/tools/goctl/util/file"
	"strings"
)

const (
	canalSvcTemplate = `package svc

import (
	"context"
	"github.com/xiaoshouchen/go-zero/zrpc"

	"hbb_micro/service/{{.ctxName}}/mq/internal/config"
	"hbb_micro/service/{{.ctxName}}/rpc/{{.ctxName}}"
)

type ServiceContext struct {
	Ctx     context.Context
	Config  config.Config
	{{.upper}}Rpc {{.ctxName}}.{{.upper}}
}

func NewServiceContext(ctx context.Context, c config.Config) *ServiceContext {
	return &ServiceContext{
		Ctx:     ctx,
		Config:  c,
		{{.upper}}Rpc: user.NewUser(zrpc.MustNewClient(c.{{.upper}}RPC)),
	}
}

`
)

func GenerateSvc(ctx *Context) error {
	// 循环替换
	t := canalSvcTemplate
	t = strings.Replace(t, "{{.ctxName}}", ctx.Name, -1)
	t = strings.Replace(t, "{{.upper}}", ctx.Upper, -1)
	routerFilePath := ctx.Output + "/" + ctx.Name + "/mq/internal/svc/service_context.go"
	err := file.GenerateFile(routerFilePath, t, false)
	if err != nil {
		return err
	}
	return nil
}
