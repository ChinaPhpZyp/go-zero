package generate

import (
	"github.com/xiaoshouchen/go-zero/tools/goctl/util/file"
	"strings"
)

const (
	canalConfigTemplate = `package config

import (
	"github.com/xiaoshouchen/go-zero/zrpc"
)

type Config struct {
	// 根据模块生成对应的rpc
	{{.upper}}RPC zrpc.RpcClientConf
}
`
)

func GenerateConfig(ctx *Context) error {
	// 循环替换
	t := canalConfigTemplate
	t = strings.Replace(t, "{{.upper}}", ctx.Upper, -1)
	routerFilePath := ctx.Output + "/" + ctx.Name + "/mq/internal/config/conf.go"
	err := file.GenerateFile(routerFilePath, t, false)
	if err != nil {
		return err
	}
	return nil
}
