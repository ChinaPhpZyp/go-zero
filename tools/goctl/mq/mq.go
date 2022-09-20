package mq

import (
	"errors"
	"fmt"
	"github.com/xiaoshouchen/go-zero/tools/goctl/util/name"
	"path/filepath"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
	"github.com/xiaoshouchen/go-zero/tools/goctl/mq/generate"
)

// Action provides the entry for goctl mongo code generation.
func Action(ctx *cli.Context) error {
	o := strings.TrimSpace(ctx.String("dir"))
	n := ctx.String("name")
	if len(n) == 0 {
		return errors.New("名称必须填")
	}
	mqName := ctx.String("mq")
	if len(mqName) == 0 {
		return errors.New("队列名必须填")
	}
	a, err := filepath.Abs(o)
	if err != nil {
		return err
	}
	upper, _, _ := name.FormFuncName(n)
	mqNameUpper, _, _ := name.FormFuncName(mqName)
	err = generate.Do(&generate.Context{
		Output:          a,
		Name:            n,
		Upper:           upper,
		ModuleName:      mqName,
		ModuleUpperName: mqNameUpper,
	})

	fmt.Println(aurora.Green("Done."))
	return err
}
