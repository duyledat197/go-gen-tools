{{define "retrieve"}}
func (d *{{.CamelCase}}Delivery) Get{{.PascalCase}}ByID(ctx context.Context, req *pb.Get{{.PascalCase}}ByIDRequest) (*pb.Get{{.PascalCase}}ByIDResponse, error) {
	{{.CamelCase}}, err := d.{{.CamelCase}}Service.Get{{.PascalCase}}ByID(ctx, id)
	return &pb.Get{{.PascalCase}}ByIDResponse{
		{{.PascalCase}}: transform.{{.PascalCase}}ToPbPtr({{.CamelCase}}),
	}, nil
}

{{end}}