package generate

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/xiaoshouchen/go-zero/tools/goctl/util/pathx"
)

const (
	category          = "mq"
	modelTemplateFile = "mq.tpl"
)

var templates = map[string]string{
	modelTemplateFile: mqMainTemplate,
}

// Category returns the mongo category.
func Category() string {
	return category
}

// Clean cleans the mongo templates.
func Clean() error {
	return pathx.Clean(category)
}

// Templates initializes the mongo templates.
func Templates(_ *cli.Context) error {
	return pathx.InitTemplates(category, templates)
}

// RevertTemplate reverts the given template.
func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}

	return pathx.CreateTemplate(category, name, content)
}

// Update cleans and updates the templates.
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return pathx.InitTemplates(category, templates)
}
