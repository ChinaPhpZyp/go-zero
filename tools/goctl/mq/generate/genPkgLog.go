package generate

import "github.com/xiaoshouchen/go-zero/tools/goctl/util/file"

const (
	PkgLogtemplate = `package queue

type Log struct {
	Messages map[string]string
}

func (l *Log) Info(key, value string) {
	l.Messages[key] = l.Messages[key] + "|" + value
}

func (l *Log) Error(value string) {
	l.Messages["error"] = l.Messages["error"] + "|" + value

}

`
)

func GeneratePkgLog(ctx *Context) error {
	routerFilePath := ctx.Output + "/../common/queue/log.go"
	err := file.GenerateFile(routerFilePath, PkgLogtemplate, false)
	if err != nil {
		return err
	}
	return nil
}
