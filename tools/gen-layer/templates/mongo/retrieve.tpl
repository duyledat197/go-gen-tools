{{define "retrieve"}}
func (r *{{.CamelCase}}Repository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.{{.PascalCase}}, error) {
	opt := &options.FindOneOptions{}
	result := &models.{{.CamelCase}}{}
	if err := r.coll.FindOne(ctx, &models.{{.CamelCase}}{ID: id}, opt).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
{{end}}