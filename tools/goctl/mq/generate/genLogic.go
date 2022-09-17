package generate

import (
	"github.com/xiaoshouchen/go-zero/tools/goctl/util/file"
	"strings"
)

const (
	mqLogicTemplate = `package logic

import (
	"context"
	"codeup.aliyun.com/huabaobao/mq/src/log"

	"hbb_micro/service/{{.ctxName}}/mq/internal/svc"
)

type {{.ModuleUpperName}}Logic struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func New{{.ModuleUpperName}}Logic(ctx context.Context, svc *svc.ServiceContext) {{.ModuleUpperName}}Logic {
	return {{.ModuleUpperName}}Logic{ctx: ctx, svc: svc}
}

// {{.ModuleUpperName}} 
func (l {{.ModuleUpperName}}Logic) Handle(message queue.Message, log queue.Log) error {
	return nil
}

`
)

func GenerateLogic(ctx *Context) error {
	t := mqLogicTemplate
	t = strings.Replace(t, "{{.ctxName}}", ctx.Name, -1)
	t = strings.Replace(t, "{{.ModuleUpperName}}", ctx.ModuleUpperName, -1)
	routerFilePath := ctx.Output + "/" + ctx.Name + "/mq/internal/logic/" + ctx.ModuleName + ".go"
	err := file.GenerateFile(routerFilePath, t, false)
	if err != nil {
		return err
	}
	return nil
}
