{{define "default"}}
package repositories

import (
	"context"

	"{{.Module}}/internal/models"
	"{{.Module}}/internal/repositories"
)

type {{.PascalCase}}Repository interface {
	{{if .IsCreate}} Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {{end}}
	{{if .IsRetrieve}} GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.{{.PascalCase}}, error) {{end}}
	{{if .IsList}} GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.{{.PascalCase}}, error) {{end}}
	{{if .IsUpdate}} Update(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}, opts ...repositories.Options) error {{end}}
	{{if .IsDelete}} Delete(ctx context.Context, id string, opts ...repositories.Options) error {{end}}
}

{{end}}