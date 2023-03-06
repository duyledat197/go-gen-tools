{{define "retrieve"}}
func (r *{{.CamelCase}}Repository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.{{.PascalCase}}, error) {
		q := models.New(r.db)
	if len(opts) > 0 && opts[0].Tx != nil {
		q = q.WithTx(opts[0].Tx)
	}
	result, err := q.Find{{.PascalCase}}ByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
{{end}}