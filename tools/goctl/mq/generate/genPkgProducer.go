package generate

import "github.com/xiaoshouchen/go-zero/tools/goctl/util/file"

const (
	PkgProducertemplate = `package queue

import (
	"encoding/json"
	"hbb_micro/common/redis"
	"os"
)

// Enqueue 消息入队
func Enqueue(scriptName string, data interface{}) bool {
	dataJson, err := json.Marshal(data)
	if err != nil {
		return false
	}
	_, err = redis.GetOneRedisClient().Execute("XADD", os.Getenv("QUEUE_STEAM_NAME"), "MAXLEN", 500000, "*", scriptName, dataJson)
	if err != nil {
		return false
	}
	return true
}

`
)

func GeneratePkgProducer(ctx *Context) error {
	routerFilePath := ctx.Output + "/../common/queue/producer.go"
	err := file.GenerateFile(routerFilePath, PkgProducertemplate, false)
	if err != nil {
		return err
	}
	return nil
}
