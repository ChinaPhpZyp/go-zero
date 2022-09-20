package generate

import (
	"github.com/xiaoshouchen/go-zero/tools/goctl/util/file"
	"strings"
)

const (
	mqMainTemplate = `package main

import (
	"codeup.aliyun.com/huabaobao/cache/drivers"
	"codeup.aliyun.com/huabaobao/mq"
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/xiaoshouchen/go-zero/core/conf"
	"hbb_micro/common/db"
	"hbb_micro/service/{{.ctxName}}/mq/internal/config"
	"hbb_micro/service/{{.ctxName}}/mq/internal/router"
	"hbb_micro/service/{{.ctxName}}/mq/internal/svc"
	"log"
	"os"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")
var envPath = flag.String("e", "../.env", "the config file")

func main() {
	err := godotenv.Load(*envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var cfg config.Config
	conf.MustLoad(*configFile, &cfg)
	ctx := svc.NewServiceContext(context.Background(), cfg)
	drivers.Init()

	eng := mq.GetInstance()
	engine := eng.New()
	router.RegisterRouter(engine, context.Background(), ctx)
	err = eng.Run(engine)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
`
)

func GenerateMQ(ctx *Context) error {
	// 循环替换
	t := mqMainTemplate
	t = strings.Replace(t, "{{.ctxName}}", ctx.Name, -1)
	routerFilePath := ctx.Output + "/" + ctx.Name + "/mq/mq.go"
	err := file.GenerateFile(routerFilePath, t, false)
	if err != nil {
		return err
	}
	return nil
}
