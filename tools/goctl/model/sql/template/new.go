package template

// New defines the template for creating model instance.
const New = `
func New{{.upperStartCamelObject}}Model(db *gorm.DB) *{{.upperStartCamelObject}}Model {
	return &{{.upperStartCamelObject}}Model{db: db}
}

func (m {{.upperStartCamelObject}}Model) baseQuery() *gorm.DB {
	return m.db.Model(&{{.upperStartCamelObject}}{})
}
`
