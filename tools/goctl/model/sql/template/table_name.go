package template

// TableName defines a template that generate the tableName method.
const TableName = `
func (m {{.upperStartCamelObject}}) TableName() string {
	return "{{.tableName}}"
}
`
