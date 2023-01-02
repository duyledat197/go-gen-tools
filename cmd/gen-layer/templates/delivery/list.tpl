{{define "list"}}
func (d *{{.CamelCase}}Delivery) GetList(ctx context.Context, req *pb.GetListRequest) (*pb.GetListResponse, error) {

	{{.CamelCase}}s, total, err := d.{{.CamelCase}}Service.GetList{{.PascalCase}}(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	result := make([]*pb.{{.PascalCase}}, 0, len({{.CamelCase}}s))
	for _, {{.CamelCase}} := range {{.CamelCase}}s {
		result = append(result, to{{.PascalCase}}Pb({{.CamelCase}}))
	}
	return &pb.GetListResponse{
		Data:  result,
		Total: int32(total),
	}, nil
}

{{end}}