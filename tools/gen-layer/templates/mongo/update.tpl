{{define "update"}}
func (r *{{.CamelCase}}Repository) Update(ctx context.Context, filter, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {
	opt := &options.UpdateOptions{}
	result, err := r.coll.UpdateMany(ctx, filter, primitive.M{
		"set": {{.CamelCase}},
	}, opt)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("update not effected")
	}
	return nil
}
{{end}}