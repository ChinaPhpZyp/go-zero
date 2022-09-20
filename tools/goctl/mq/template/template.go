package template

// Text provides the default template for model to generate
var Text = `package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/joho/godotenv"
	pbe "github.com/withlin/canal-go/protocol/entry"
	"hbb_micro/common/canal"
	"hbb_micro/service/user/canal/router"
	"os"
	"strconv"
	"time"
)

func main() {
	godotenv.Load("../.env")
	canal.Init()
	router.RegisterCanalRouter()

	start()
}

func start() {
	batchSize, _ := strconv.ParseInt(os.Getenv("CANAL_BATCH_SIZE"), 10, 32)
	for {
		message, err := canal.CanalClient.Get(int32(batchSize), nil, nil)
		fmt.Println(message)
		if err != nil {
			fmt.Println("获取消息失败", err)
			time.Sleep(300 * time.Millisecond)
			continue
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			continue
		}
		for _, entry := range message.Entries {
			if entry.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
				continue
			}
			rowChange := new(pbe.RowChange)
			err = proto.Unmarshal(entry.GetStoreValue(), rowChange)
			checkError(err)
			if rowChange != nil {
				eventType := rowChange.GetEventType()
				header := entry.GetHeader()
				for _, rowData := range rowChange.GetRowDatas() {
					var eventTypeName string
					switch eventType {
					case pbe.EventType_DELETE:
						eventTypeName = "delete"
					case pbe.EventType_INSERT:
						eventTypeName = "insert"
					case pbe.EventType_UPDATE:
						eventTypeName = "update"
					}
					tableName := header.GetTableName()
					fmt.Println(tableName)
					if execFunc, ok := router.CanalRouter[tableName+"@"+eventTypeName]; ok {
						fmt.Println("执行" + tableName + "@" + eventTypeName)
						go execFunc(rowData.GetAfterColumns())
					}
				}
			}
		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

`
