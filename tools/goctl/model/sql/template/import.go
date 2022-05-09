package template

const (
	// Imports defines a import template for model in cache case
	Imports = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/xiaoshouchen/go-zero/core/stores/builder"
	"github.com/xiaoshouchen/go-zero/core/stores/cache"
	"github.com/xiaoshouchen/go-zero/core/stores/sqlc"
	"github.com/xiaoshouchen/go-zero/core/stores/sqlx"
	"github.com/xiaoshouchen/go-zero/core/stringx"
)
`
	// ImportsNoCache defines a import template for model in normal case
	ImportsNoCache = `import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}

	"github.com/xiaoshouchen/go-zero/core/stores/builder"
	"github.com/xiaoshouchen/go-zero/core/stores/sqlc"
	"github.com/xiaoshouchen/go-zero/core/stores/sqlx"
	"github.com/xiaoshouchen/go-zero/core/stringx"
)
`
)
