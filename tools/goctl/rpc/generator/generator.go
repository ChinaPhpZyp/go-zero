package generator

import (
	"log"

	conf "github.com/xiaoshouchen/go-zero/tools/goctl/config"
	"github.com/xiaoshouchen/go-zero/tools/goctl/env"
	"github.com/xiaoshouchen/go-zero/tools/goctl/util/console"
)

// Generator defines the environment needs of rpc service generation
type Generator struct {
	log     console.Console
	cfg     *conf.Config
	verbose bool
}

// NewGenerator returns an instance of Generator
func NewGenerator(style string, verbose bool) *Generator {
	cfg, err := conf.NewConfig(style)
	if err != nil {
		log.Fatalln(err)
	}
	log := console.NewColorConsole(verbose)
	return &Generator{
		log:     log,
		cfg:     cfg,
		verbose: verbose,
	}
}

// Prepare provides environment detection generated by rpc service,
// including go environment, protoc, whether protoc-gen-go is installed or not
func (g *Generator) Prepare() error {
	return env.Prepare(true, true, g.verbose)
}
