package template

// Types defines a template for types in model.
const Types = `
type (
	{{.upperStartCamelObject}}Model struct {
		db *gorm.DB
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
`
