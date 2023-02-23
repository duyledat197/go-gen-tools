package main

import (
	"fmt"
)

type DBType string

const (
	Mongo    DBType = "Mongo"
	Postgres DBType = "Postgres"
)

type Database struct {
	Host     string
	Port     string
	Database string
	UserName string
	Password string
	SSLMode  string

	Type DBType
}

func (p *Database) GetConnectionString() string {
	var prefix string
	switch p.Type {
	case Mongo:
		prefix = "mongo+srv"
	case Postgres:
		prefix = "postgres"
	}

	return fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=%s", prefix, p.UserName, p.Password, p.Host, p.Database, p.SSLMode)
}

type ConnectionAddr struct {
	Host string
	Port string
}

type Config struct {
	// database
	PostgresDB *Database
	MongoDB    *Database

	// connection address
	HTTP *ConnectionAddr
	GRPC *ConnectionAddr
}
