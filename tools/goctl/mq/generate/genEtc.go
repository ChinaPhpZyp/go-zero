package generate

import (
	"github.com/xiaoshouchen/go-zero/tools/goctl/util/file"
	"strings"
)

const (
	canalEtcTemplate = `Name: mq
{{.upper}}RPC:
 Etcd:
  Hosts:
   - 127.0.0.1:2379
  Key: {{.ctxName}}.rpc
`
)

func GenerateEtc(ctx *Context) error {
	// 循环替换
	t := canalEtcTemplate
	t = strings.Replace(t, "{{.upper}}", ctx.Upper, -1)
	t = strings.Replace(t, "{{.ctxName}}", ctx.Name, -1)
	routerFilePath := ctx.Output + "/" + ctx.Name + "/mq/etc/" + ctx.Name + ".yaml"
	err := file.GenerateFile(routerFilePath, t, false)
	if err != nil {
		return err
	}
	return nil
}
