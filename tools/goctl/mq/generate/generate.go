package generate

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/xiaoshouchen/go-zero/core/logx"
)

// Context defines the model generation data what they needs
type Context struct {
	Output          string
	Name            string
	Upper           string
	During          string
	Desc            string
	ModuleName      string
	ModuleUpperName string
}

// Do executes model template and output the result into the specified file path
func Do(ctx *Context) error {
	// 生成pkg
	logx.Must(GeneratePkgConsumer(ctx))
	logx.Must(GeneratePkgEngine(ctx))
	logx.Must(GeneratePkgLog(ctx))
	logx.Must(GeneratePkgMessage(ctx))
	logx.Must(GeneratePkgProducer(ctx))
	// 创建ETC
	logx.Must(GenerateEtc(ctx))
	// 生成config
	logx.Must(GenerateConfig(ctx))
	// 生成svc
	logx.Must(GenerateSvc(ctx))
	// 生成logic
	logx.Must(GenerateLogic(ctx))
	// 创建路由
	logx.Must(GenerateRouter(ctx))
	// 创建启动文件
	logx.Must(GenerateMQ(ctx))
	fmt.Println(aurora.Green("Done."))
	return nil
}
