{{define "default"}}
package services

import (
	"context"
	"time"

	"{{.Module}}/internal/models"
	"{{.Module}}/internal/repositories"
)

type {{.PascalCase}}Service interface {
	Create{{.PascalCase}}(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}) error
	Get{{.PascalCase}}ByID(ctx context.Context, id string) (*models.{{.PascalCase}}, error)
	GetList{{.PascalCase}}(ctx context.Context, offset, limit int) ([]*models.{{.PascalCase}}, error)
	Update{{.PascalCase}}(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}) error
	Delete{{.PascalCase}}(ctx context.Context, id string) error
}

type {{.CamelCase}}Service struct {
	{{.CamelCase}}Repo repositories.{{.PascalCase}}Repository
}

func New{{.PascalCase}}Service({{.CamelCase}}Repo repositories.{{.PascalCase}}Repository) {{.PascalCase}}Service {
	return &{{.CamelCase}}Service{
		{{.CamelCase}}Repo: {{.CamelCase}}Repo,
	}
}
{{end}}