{{define "delete"}}
func (r *{{.CamelCase}}Repository) Delete(ctx context.Context, id string, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}
{{end}}