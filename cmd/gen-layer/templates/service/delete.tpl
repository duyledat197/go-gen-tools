{{define "delete"}}
func (s *{{.CamelCase}}Service) Delete(ctx context.Context, id string) error {
	if err := s.{{.CamelCase}}Repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
{{end}}