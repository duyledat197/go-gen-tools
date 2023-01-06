{{define "delete"}}
func (u *{{.CamelCase}}Repository) Delete(ctx context.Context, id string, opts ...repositories.Options) error {
	return nil
}
{{end}}