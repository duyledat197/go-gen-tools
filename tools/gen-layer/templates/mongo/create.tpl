{{define "create"}}
func (r *{{.CamelCase}}Repository) Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {
	if _, err := r.coll.InsertOne(ctx, {{.CamelCase}}, &options.InsertOneOptions{}); err != nil {
		return err
	}
	return nil
}
{{end}}