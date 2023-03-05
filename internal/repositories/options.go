package repositories

import (
	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type Options struct {
	// transaction for DBMS
	Tx pgx.Tx

	// transaction for mongodb
	SessionContext mongo.SessionContext
}
