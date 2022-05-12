package template

import (
	"fmt"

	"github.com/xiaoshouchen/go-zero/tools/goctl/util"
)

// ModelCustom defines a template for extension
const ModelCustom = `package {{.pkg}}

`

// ModelGen defines a template for model
var ModelGen = fmt.Sprintf(`%s

package {{.pkg}}
{{.imports}}
{{.vars}}
{{.types}}
{{.new}}
{{.insert}}
{{.find}}
{{.update}}
{{.delete}}
{{.extraMethod}}
{{.tableName}}
`, util.DoNotEditHead)
