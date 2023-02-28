package config

import (
	"fmt"
	"time"
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

	MaxConnection int
	Timeout       time.Duration
	Type          DBType
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

func (p *ConnectionAddr) GetConnectionString() string {
	return fmt.Sprintf("%s:%s", p.Host, p.Port)
}

type Config struct {
	ServiceName string
	// database
	PostgresDB *Database
	MongoDB    *Database
	ElasticDB  *Database
	RedisDB    *Database

	// connection address
	HTTP *ConnectionAddr
	GRPC *ConnectionAddr

	Consul     *ConnectionAddr
	Tracer     *ConnectionAddr
	Prometheus *ConnectionAddr
	Kafka      *ConnectionAddr
}
