{{define "list"}}
func (d *{{.CamelCase}}Delivery) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {
	{{.CamelCase}}s, err := d.{{.CamelCase}}Service.GetList(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	return &pb.GetListResponse{
		Data: transform.{{.PascalCase}}ToPbPtrList({{.CamelCase}}s),
	}, nil
}
{{end}}