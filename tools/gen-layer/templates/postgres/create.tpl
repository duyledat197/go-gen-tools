{{define "create"}}
func (r *{{.CamelCase}}Repository) Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {
	q := models.New(r.db)
	if len(opts) > 0 && opts[0].Tx != nil {
		q = q.WithTx(opts[0].Tx)
	}
	b, err := json.Marshal({{.CamelCase}})
	if err != nil {
		return err
	}
	var params models.Create{{.PascalCase}}Params
	if err := json.Unmarshal(b, &params); err != nil {
		return err
	}
	if _, err := q.Create{{.PascalCase}}(ctx, params); err != nil {
		return err
	}
	return nil
}
{{end}}