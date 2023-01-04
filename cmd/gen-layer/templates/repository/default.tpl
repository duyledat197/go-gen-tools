{{define "default"}}
package repositories

import (
	"context"

	"{{.Module}}/internal/models"
)

type {{.PascalCase}}Repository interface {
	Create(ctx context.Context, {{.CamelCase}} *models.{{.PascalCase}}) error
	GetByID(ctx context.Context, id string) (*models.{{.PascalCase}}, error)
	GetList(ctx context.Context, offset, limit int) ([]*models.{{.PascalCase}}, error)
	Update(ctx context.Context, id string, {{.CamelCase}} *models.{{.PascalCase}}) error
	Delete(ctx context.Context, id string) error
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