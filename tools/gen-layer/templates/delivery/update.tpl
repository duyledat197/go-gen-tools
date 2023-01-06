{{define "update"}}
func (d *{{.CamelCase}}Delivery) Update{{.PascalCase}}(ctx context.Context, req *pb.Update{{.PascalCase}}Request) (*pb.Update{{.PascalCase}}Response, error) {
	if err := d.{{.CamelCase}}Service.Update(ctx, req.{{.PascalCase}}.Id, transform.PbTo{{.PascalCase}}Ptr(req.Get{{.PascalCase}}())); err != nil {
		return nil, err
	}
	return &pb.Update{{.PascalCase}}Response{
		Success: true,
	}, nil
}
{{end}}