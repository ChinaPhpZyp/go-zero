package template

// New defines the template for creating model instance.
const New = `
func New{{.upperStartCamelObject}}Model(db ...*gorm.DB) *{{.upperStartCamelObject}}Model {
	if len(db) >= 1 {
		return &{{.upperStartCamelObject}}Model{db: db[0]}
	}
	return &{{.upperStartCamelObject}}Model{}
}

func (m {{.upperStartCamelObject}}Model) baseQuery() *gorm.DB {
	if m.db != nil {
		return m.db.Model(&{{.upperStartCamelObject}}{})
	}
	return db.Db.Model(&{{.upperStartCamelObject}}{})
}
`
