package generate

import "github.com/xiaoshouchen/go-zero/tools/goctl/util/file"

const (
	PkgMessagetemplate = `package queue

import redis2 "github.com/gomodule/redigo/redis"

type Message struct {
	QueueName string
	MessageId string
	Key       string
	Value     string
}

func getMessage(res interface{}, err error) ([]Message, error) {
	var msg []Message
	valueRes, err1 := redis2.Values(res, err)
	for kIndex := 0; kIndex < len(valueRes); kIndex++ {
		var keyInfo = valueRes[kIndex].([]interface{})
		var key = string(keyInfo[0].([]byte))
		var idList = keyInfo[1].([]interface{})
		for idIndex := 0; idIndex < len(idList); idIndex++ {
			var idInfo = idList[idIndex].([]interface{})
			var id = string(idInfo[0].([]byte))
			var fieldList = idInfo[1].([]interface{})
			var field = string(fieldList[0].([]byte))
			var value = string(fieldList[1].([]byte))
			msg = append(msg, Message{key, id, field, value})
		}
	}
	return msg, err1
}

`
)

func GeneratePkgMessage(ctx *Context) error {
	routerFilePath := ctx.Output + "/../common/queue/message.go"
	err := file.GenerateFile(routerFilePath, PkgMessagetemplate, false)
	if err != nil {
		return err
	}
	return nil
}
