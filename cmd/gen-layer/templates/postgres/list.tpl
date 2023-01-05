{{define "list"}}
func (u *{{.CamelCase}}Repository) GetList(ctx context.Context, offset, limit int) ([]*models.{{.PascalCase}}, error) {
	return nil, nil
}
{{end}}