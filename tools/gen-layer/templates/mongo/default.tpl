{{define "default"}}
package mongo

import (
	"context"

	"{{.Module}}/internal/models"
    "{{.Module}}/internal/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type {{.CamelCase}}Repository struct {
	coll *mongo.Collection
}

func New{{.PascalCase}}Repository(coll *mongo.Collection) repositories.{{.PascalCase}}Repository {
	return &{{.CamelCase}}Repository{
		coll,
	}
}
{{end}}