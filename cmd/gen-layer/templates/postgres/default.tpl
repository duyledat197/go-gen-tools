{{define "default"}}
package postgres

import (
	"context"

	"{{.Module}}/internal/models"
    "{{.Module}}/internal/repositories"
)

type {{.CamelCase}}Repository struct {
	db *models.Queries
}

func New{{.PascalCase}}Repository(q *models.Queries) repositories.{{.PascalCase}}Repository {
	return &{{.CamelCase}}Repository{
		db: q,
	}
}
{{end}}