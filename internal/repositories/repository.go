package repositories

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Options struct {
	// transaction for DBMS
	Tx *sql.Tx

	// transaction for mongodb
	SessionContext mongo.SessionContext
}
