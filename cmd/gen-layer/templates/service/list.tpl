{{define "list"}}
func (s *{{.CamelCase}}Service) GetList{{.PascalCase}}(ctx context.Context, offset, limit int) ([]*models.{{.PascalCase}}, error) {
	{{.CamelCase}}s, err := s.{{.CamelCase}}Repo.GetList(ctx, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	return {{.CamelCase}}s, nil
}
{{end}}