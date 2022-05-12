package template

const (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"context"
	"database/sql"
	"fmt"
	{{if .time}}"time"{{end}}
)
`
)
