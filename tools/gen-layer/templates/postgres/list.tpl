{{define "list"}}
func (r *{{.CamelCase}}Repository) GetList(ctx context.Context, filter *models.{{.PascalCase}}, offset, limit int, opts ...repositories.Options) ([]*models.{{.PascalCase}}, error) {
	q := models.New(r.db)
	if len(opts) > 0 && opts[0].Tx != nil {
		q = q.WithTx(opts[0].Tx)
	}

	b, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}
	var params models.GetList{{.PascalCase}}Params
	if err := json.Unmarshal(b, &params); err != nil {
		return nil, err
	}

	result, err := q.GetList{{.PascalCase}}(ctx, params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
{{end}}