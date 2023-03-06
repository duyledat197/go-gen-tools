{{define "delete"}}
func (r *{{.CamelCase}}Repository) Delete(ctx context.Context, filter *models.{{.CamelCase}}, opts ...repositories.Options) error {
	opt := &options.DeleteOptions{}
	result, err := r.coll.DeleteMany(ctx, filter, opt)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("delete not effected")
	}
	return nil
}
{{end}}