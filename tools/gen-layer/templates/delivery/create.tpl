{{define "create"}}
func (d *{{.CamelCase}}Delivery) Create{{.PascalCase}}(ctx context.Context, req *pb.Create{{.PascalCase}}Request) (*pb.Create{{.PascalCase}}Response, error) {
	if err := d.{{.CamelCase}}Service.Create(ctx, transform.PbTo{{.PascalCase}}Ptr(req.Get{{.PascalCase}}())); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("Create: %v", err).Error())
	}
	return &pb.Create{{.PascalCase}}Response{}, nil
}
{{end}}