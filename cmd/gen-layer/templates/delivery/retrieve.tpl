{{define "retrieve"}}
func (d *{{.CamelCase}}Delivery) Get{{.PascalCase}}ByID(ctx context.Context, req *pb.Get{{.PascalCase}}ByIDRequest) (*pb.Get{{.PascalCase}}ByIDResponse, error) {
	{{.CamelCase}}, err := d.{{.CamelCase}}Service.GetByID(ctx, req.Get{{.PascalCase}}ID())
	if err != nil {
		return nil, err
	}
	return &pb.Get{{.PascalCase}}ByIDResponse{
		Data: transform.{{.PascalCase}}ToPbPtr({{.CamelCase}}),
	}, nil
}
{{end}}