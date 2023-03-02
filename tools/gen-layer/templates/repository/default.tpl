{{define "default"}}
package repositories

import (
	"context"

	"{{.Module}}/internal/models"
)

type {{.PascalCase}}Repository interface {
	{{if .IsCreate}} Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}, opts ...Options) error {{end}}
	{{if .IsRetrieve}} GetByID(ctx context.Context, id string, opts ...Options) (*models.{{.PascalCase}}, error) {{end}}
	{{if .IsList}} GetList(ctx context.Context, offset, limit int, opts ...Options) ([]*models.{{.PascalCase}}, error) {{end}}
	{{if .IsUpdate}} Update(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}, opts ...Options) error {{end}}
	{{if .IsDelete}} Delete(ctx context.Context, id string, opts ...Options) error {{end}}
}

{{end}}