{{define "default"}}
package deliveries

import (
	"context"
	"fmt"

	"{{.Module}}/internal/services"
	"{{.Module}}/pb"
	"{{.Module}}/transform"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type {{.CamelCase}}Delivery struct {
	{{.CamelCase}}Service services.{{.PascalCase}}Service
	pb.Unimplemented{{.PascalCase}}ServiceServer
}

func New{{.PascalCase}}Delivery({{.CamelCase}}Service services.{{.PascalCase}}Service) pb.{{.PascalCase}}ServiceServer {
	return &{{.CamelCase}}Delivery{
		{{.CamelCase}}Service: {{.CamelCase}}Service,
	}
}
{{end}}