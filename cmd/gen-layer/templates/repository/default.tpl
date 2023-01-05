{{define "default"}}
package repositories

import (
	"context"

	"{{.Module}}/internal/models"
)

type {{.PascalCase}}Repository interface {
	{{if .IsCreate}} Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}) error {{end}}
	{{if .IsRetrieve}} GetByID(ctx context.Context, id string) (*models.{{.PascalCase}}, error) {{end}}
	{{if .IsList}} GetList(ctx context.Context, offset, limit int) ([]*models.{{.PascalCase}}, error) {{end}}
	{{if .IsUpdate}} Update(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}) error {{end}}
	{{if .IsDelete}} Delete(ctx context.Context, id string) error {{end}}
}

type {{.CamelCase}}Repository struct {
	db *models.Queries
}

func New{{.PascalCase}}Repository(q *models.Queries) {{.PascalCase}}Repository {
	return &{{.CamelCase}}Repository{
		db: q,
	}
}
{{end}}