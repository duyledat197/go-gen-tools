{{define "default"}}
package services

import (
	"context"
	"time"

	"{{.Module}}/internal/models"
	"{{.Module}}/internal/repositories"
	
	"github.com/jackc/pgtype"
)

type {{.PascalCase}}Service interface {
	{{if .IsCreate}} Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}) error {{end}}
	{{if .IsRetrieve}} GetByID(ctx context.Context, id string) (*models.{{.PascalCase}}, error) {{end}}
	{{if .IsList}} GetList(ctx context.Context, offset, limit int) ([]*models.{{.PascalCase}}, error) {{end}}
	{{if .IsUpdate}} Update(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}) error {{end}}
	{{if .IsDelete}} Delete(ctx context.Context, id string) error {{end}}
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