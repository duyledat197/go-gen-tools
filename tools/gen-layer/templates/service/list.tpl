{{define "list"}}
func (s *{{.CamelCase}}Service) GetList(ctx context.Context, offset, limit int) ([]*models.{{.PascalCase}}, error) {
	{{.CamelCase}}s, err := s.{{.CamelCase}}Repo.GetList(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return {{.CamelCase}}s, nil
}
{{end}}