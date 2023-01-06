{{define "retrieve"}}
func (s *{{.CamelCase}}Service) GetByID(ctx context.Context, id string) (*models.{{.PascalCase}}, error) {
	{{.CamelCase}}, err := s.{{.CamelCase}}Repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return {{.CamelCase}}, nil
}
{{end}}