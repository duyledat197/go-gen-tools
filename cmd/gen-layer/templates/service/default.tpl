{{define "default"}}
package services

import (
	"context"
	"time"

	"{{.Module}}/internal/models"
	"{{.Module}}/internal/repositories"
)

type {{.PascalCase}}Service interface {
	Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}) error
	GetByID(ctx context.Context, id string) (*models.{{.PascalCase}}, error)
	GetList(ctx context.Context, offset, limit int) ([]*models.{{.PascalCase}}, error)
	Update(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}) error
	Delete(ctx context.Context, id string) error
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