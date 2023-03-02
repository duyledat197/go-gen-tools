{{define "update"}}
func (r *{{.CamelCase}}Repository) Update(ctx context.Context, filter *models.{{.PascalCase}}, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {
	result, err := r.coll.UpdateMany(ctx, filter, primitive.M{
		"set": {{.CamelCase}},
	}, &options.UpdateOptions{})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("update not effected")
	}
	return nil
}
{{end}}