package main

import (
    "codeup.aliyun.com/huabaobao/cache/drivers"
	"github.com/joho/godotenv"

	"flag"
	"fmt"

	{{.importPackages}}
)

var configFile = flag.String("f", "etc/{{.serviceName}}.yaml", "the config file")
var envPath = flag.String("e", "../.env", "the env file")

func main() {
	flag.Parse()
	_ = godotenv.Load(*envPath)

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)

	auth := middleware.NewAuth(c.Auth.AccessSecret)
	server.Use(middleware.Close)    //停服
	server.Use(auth.Check)          //JWT验证
	server.Use(middleware.IpBlock)  //预防大量DDOS攻击
	server.Use(middleware.Recovery) //致命错误提醒
	server.Use(middleware.Throttle) //限流

	defer server.Stop()

	drivers.Init()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
