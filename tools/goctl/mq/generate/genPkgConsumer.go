package generate

import "github.com/xiaoshouchen/go-zero/tools/goctl/util/file"

const (
	PkgConsumertemplate = `package queue

import (
	"encoding/json"
	"fmt"
	"hbb_micro/common/redis"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type MqStatus struct {
	mu    sync.Mutex
	isRun bool
}

func (r *MqStatus) isRunning() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.isRun
}

func (r *MqStatus) ChangeRunning(run bool) {
	r.mu.Lock()
	r.isRun = run
	r.mu.Unlock()
}

//Run 消费者模块
func Run(engine *Engine, steam, group string) error {
	status := &MqStatus{isRun: true}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-quit
		status.ChangeRunning(false)
	}()
	//最多支持100个协程同时执行
	max := make(chan int, 100)
	for {
		// 如果收到k8s的term信号，则停止运行
		if !status.isRunning() {
			break
		}
		_, _ = redis.GetOneRedisClient().Execute("XGROUP", "CREATE", steam, group, "0", "MKSTREAM")
		max <- 1
		// 用于创建消费者
		consumeName := fmt.Sprintf("consume_%d_%d", time.Now().UnixNano(), rand.Int())
		// 监听对应的消费组，并使用一个消费者去消费某个队列
		res, readErr := redis.GetOneRedisClient().Execute("XREADGROUP", "GROUP", group, consumeName,
			"COUNT", 1, "BLOCK", 10000, "STREAMS", steam, ">")
		messages, msgErr := getMessage(res, readErr)
		if msgErr != nil {
			fmt.Println(msgErr.Error())
			continue
		}
		value := messages[0]
		// 此处由原先的[]byte转成Message格式，用于后续做任务确认
		go func(params Message) {
			defer func() {
				if pErr := recover(); pErr != nil {

				}
			}()
			startTime := time.Now().UnixNano()
			var log2 = Log{Messages: make(map[string]string)}
			log2.Info("start_time", fmt.Sprintf("%d", startTime))
			defer func(map[string]string) {
				// 记录信息
				jsonString, _ := json.Marshal(log2.Messages)
				fmt.Println(jsonString)
				// 记录结束时间
				endTime := time.Now().UnixNano()
				log2.Info("end_time", fmt.Sprintf("%d", endTime))
			}(log2.Messages)
			// 执行
			err := engine.Exec(params, log2)
			if err != nil {
				log2.Error(err.Error())
			}
			//执行完释放
			<-max
		}(value)

	}
	time.Sleep(time.Minute * 5)
	return nil
}
`
)

func GeneratePkgConsumer(ctx *Context) error {
	routerFilePath := ctx.Output + "/../common/queue/consumer.go"
	err := file.GenerateFile(routerFilePath, PkgConsumertemplate, false)
	if err != nil {
		return err
	}
	return nil
}
