{{define "create"}}
func (r *{{.CamelCase}}Repository) Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {
	opt := &options.InsertOneOptions{}
	if _, err := r.coll.InsertOne(ctx, {{.CamelCase}}, opt); err != nil {
		return err
	}
	return nil
}
{{end}}