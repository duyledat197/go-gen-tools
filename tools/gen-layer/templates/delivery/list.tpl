{{define "list"}}
func (d *{{.CamelCase}}Delivery) GetList(ctx context.Context, req *pb.GetList{{.PascalCase}}Request) (*pb.GetList{{.PascalCase}}Response, error) {
	{{.CamelCase}}s, err := d.{{.CamelCase}}Service.GetList(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	return &pb.GetList{{.PascalCase}}Response{
		Data: transform.{{.PascalCase}}ToPbPtrList({{.CamelCase}}s),
	}, nil
}
{{end}}