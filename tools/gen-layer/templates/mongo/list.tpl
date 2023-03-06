{{define "list"}}
func (r *{{.CamelCase}}Repository) GetList(ctx context.Context, filter *models.{{.CamelCase}}, offset, limit int, opts ...repositories.Options) ([]*models.{{.PascalCase}}, error) {
	opt := &options.FindOptions{}
	var result []*models.{{.CamelCase}}
	cur, err := r.coll.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
{{end}}