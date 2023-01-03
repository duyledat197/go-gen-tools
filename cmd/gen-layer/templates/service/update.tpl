{{define "update"}}
func (s *{{.CamelCase}}Service) Update{{.PascalCase}}(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}) error {
	{{.CamelCase}}.UpdatedAt = time.Now()
	if err := s.{{.CamelCase}}Repo.Update(ctx, id, {{.CamelCase}}); err != nil {
		return err
	}
	return nil
}
{{end}}