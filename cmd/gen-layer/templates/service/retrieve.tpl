{{define "retrieve"}}
func (s *{{.CamelCase}}Service) Get{{.PascalCase}}ByID(ctx context.Context, id string) (*models.{{.PascalCase}}, error) {
	{{.CamelCase}}, err := s.{{.CamelCase}}Repo.FindBy{{.PascalCase}}ID(ctx, id)
	if err != nil {
		return nil, err
	}
	return {{.CamelCase}}, nil
}
{{end}}