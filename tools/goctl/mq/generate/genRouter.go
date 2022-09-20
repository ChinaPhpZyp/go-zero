package generate

import (
	"github.com/xiaoshouchen/go-zero/tools/goctl/util/file"
	"io/ioutil"
	"os"
	"strings"
)

const (
	routerNewtemplate = `package router

import (
	"codeup.aliyun.com/huabaobao/mq/src/log"
	"context"

	"hbb_micro/service/{{.ctxName}}/mq/internal/logic"
	"hbb_micro/service/{{.ctxName}}/mq/internal/svc"
)

func RegisterRouter(engine *queue.Engine, ctx context.Context, svc *svc.ServiceContext) {
	{{.routerInsertTemplate}}
}`

	routerInsertTemplate = `
	{{.modelUpperName}}Logic := logic.New{{.modelUpperName}}Logic(ctx, svc)
	engine.Set("{{.modelName}}", {{.modelUpperName}}Logic.{{.modelUpperName}})`
)

func GenerateRouter(ctx *Context) error {
	routerFilePath := ctx.Output + "/" + ctx.Name + "/mq/internal/router/router.go"
	f, err := os.OpenFile(routerFilePath, os.O_RDWR, 700)
	// 判断文件是否存在
	if err != nil && !os.IsExist(err) {
		// 文件不存在
		t := routerNewtemplate
		t = strings.Replace(t, "{{.ctxName}}", ctx.Name, -1)
		t1 := routerInsertTemplate
		t1 = strings.Replace(t1, "{{.modelUpperName}}", ctx.ModuleUpperName, -1)
		t1 = strings.Replace(t1, "{{.modelName}}", ctx.ModuleName, -1)
		t = strings.Replace(t, "{{.routerInsertTemplate}}", t1, -1)
		return file.GenerateFile(routerFilePath, t, true)
	} else {
		// 获得内容信息
		readAll, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		readAll = readAll[:len(readAll)-1]
		readString := strings.TrimRight(string(readAll), "}")
		t1 := routerInsertTemplate
		t1 = strings.Replace(t1, "{{.modelName}}", ctx.ModuleName, -1)
		t1 = strings.Replace(t1, "{{.modelUpperName}}", ctx.ModuleUpperName, -1)
		readString = readString + t1 + "\n" + "}"
		f.Close()
		f1, err := os.OpenFile(routerFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 700)
		f1.Write([]byte(readString))
		f1.Close()
		return nil
	}
}
