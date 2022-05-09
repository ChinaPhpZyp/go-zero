package main

import (
	"github.com/xiaoshouchen/go-zero/core/load"
	"github.com/xiaoshouchen/go-zero/core/logx"
	"github.com/xiaoshouchen/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
