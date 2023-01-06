{{define "delete"}}
func (d *{{.CamelCase}}Delivery) Delete{{.PascalCase}}(ctx context.Context, req *pb.Delete{{.PascalCase}}Request) (*pb.Delete{{.PascalCase}}Response, error) {
	if err := d.{{.CamelCase}}Service.Delete(ctx, req.{{.PascalCase}}.Id); err != nil {
		return nil, err
	}
	return &pb.Delete{{.PascalCase}}Response{
		Success: true,
	}, nil
}
{{end}}