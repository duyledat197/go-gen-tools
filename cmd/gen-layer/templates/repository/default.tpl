{{define "default"}}
package deliveries

import (
	"context"
	"fmt"
	"time"

	"{{.Module}}/internal/models"
	"{{.Module}}/internal/services"
	"{{.Module}}/pb"
)

type {{.PascalCase}}Delivery struct {
	{{.CamelCase}}Service services.{{.PascalCase}}Service
	pb.Unimplemented{{.PascalCase}}ServiceServer
}

func New{{.PascalCase}}Delivery({{.CamelCase}}Service services.{{.PascalCase}}Service) pb.{{.PascalCase}}ServiceServer {
	return &{{.CamelCase}}Delivery{
		{{.CamelCase}}Service: {{.CamelCase}}Service,
	}
}

{{end}}