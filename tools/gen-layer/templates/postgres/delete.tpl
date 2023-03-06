{{define "delete"}}
func (r *{{.CamelCase}}Repository) Delete(ctx context.Context, filter *models.{{.PascalCase}}, opts ...repositories.Options) error {
	q := models.New(r.db)
	if len(opts) > 0 && opts[0].Tx != nil {
		q = q.WithTx(opts[0].Tx)
	}
	b, err := json.Marshal(filter)
	if err != nil {
		return err
	}
	var params models.Delete{{.PascalCase}}Params
	if err := json.Unmarshal(b, &params); err != nil {
		return err
	}
	if _, err := q.Delete{{.PascalCase}}(ctx, params); err != nil {
		return err
	}
	return nil
}
{{end}}