{{define "delete"}}
func (u *{{.CamelCase}}Repository) Delete(ctx context.Context, id string) error {
	return nil
}
{{end}}