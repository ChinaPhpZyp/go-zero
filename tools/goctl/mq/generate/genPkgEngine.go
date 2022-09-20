package generate

import "github.com/xiaoshouchen/go-zero/tools/goctl/util/file"

const (
	PkgEnginetemplate = `package queue

import "errors"

type Engine struct {
	routerMap map[string]func(data Message, log Log) error
}

func NewEngine() *Engine {
	return &Engine{}
}

func (r *Engine) Set(str string, f func(data Message, log Log) error) {
	if nil == r.routerMap {
		r.routerMap = make(map[string]func(data Message, log Log) error)
	}
	r.routerMap[str] = f
}

func (r *Engine) Exec(message Message, log Log) error {
	if funcName, ok := r.routerMap[message.Key]; ok {
		return funcName(message, log)
	}
	return errors.New("函数不存在")
}
`
)

func GeneratePkgEngine(ctx *Context) error {
	routerFilePath := ctx.Output + "/../common/queue/engine.go"
	err := file.GenerateFile(routerFilePath, PkgEnginetemplate, false)
	if err != nil {
		return err
	}
	return nil
}
