package main

import (
	"flag"
	"fmt"
	"hbb_micro/common/db"

	{{.imports}}
    "codeup.aliyun.com/huabaobao/cache/drivers"
	"github.com/joho/godotenv"
	"github.com/xiaoshouchen/go-zero/core/conf"
	"github.com/xiaoshouchen/go-zero/core/service"
	"github.com/xiaoshouchen/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")
var envPath = flag.String("e", "../.env", "the env file")

func main() {
	flag.Parse()
    _ = godotenv.Load(*envPath)

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.New{{.serviceNew}}Server(ctx)

    db.Init()
   	drivers.Init()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		{{.pkg}}.Register{{.service}}Server(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
