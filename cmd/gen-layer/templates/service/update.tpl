{{define "update"}}
func (s *{{.CamelCase}}Service) Update(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}) error {
	{{.CamelCase}}.UpdatedAt = pgtype.Timestamptz{
		Time: time.Now(),
	}
	if err := s.{{.CamelCase}}Repo.Update(ctx, id, {{.CamelCase}}); err != nil {
		return err
	}
	return nil
}
{{end}}