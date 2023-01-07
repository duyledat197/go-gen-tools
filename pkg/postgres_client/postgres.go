package postgres_client

import (
	"context"
	"fmt"
	"net/url"
	"time"

	pgxzap "github.com/jackc/pgx-zap"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	pgxxray "github.com/jackhopner/pgx-xray-tracer"
	"go.uber.org/zap"
)

type Options struct {
	//* enable aws xray
	IsEnableXray bool

	//* enable log trace
	IsEnableLogTrace bool
}

type PostgresClient struct {
	Pool          *pgxpool.Pool
	ConnectionURI string
	MaxConnection int32
	Config        *pgxpool.Config

	Logger  *zap.Logger
	Options *Options
}

func (c *PostgresClient) Init(ctx context.Context) *PostgresClient {
	u, err := url.Parse(c.ConnectionURI)
	if err != nil {
		c.Logger.Panic("cannot create new connection to Postgres (failed to parse URI)", zap.Error(err))
	}
	config, err := pgxpool.ParseConfig(c.ConnectionURI)
	if err != nil {
		c.Logger.Panic("cannot read PG_CONNECTION_URI", zap.Error(err))
	}

	config.MaxConns = c.MaxConnection
	config.MaxConnIdleTime = 15 * time.Second
	config.HealthCheckPeriod = 600 * time.Millisecond

	if c.Options != nil {
		options := c.Options
		switch {
		case options.IsEnableLogTrace:
			tracer := &tracelog.TraceLog{
				Logger:   pgxzap.NewLogger(c.Logger),
				LogLevel: tracelog.LogLevelTrace,
			}
			config.ConnConfig.Tracer = tracer
			break
		case options.IsEnableXray:
			config.ConnConfig.Tracer = pgxxray.NewPGXTracer()
			break
		}
	}

	c.Logger.Info("NewConnectionPool max connection", zap.Int32("max connection", c.MaxConnection))

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		c.Logger.Panic(fmt.Sprintf("cannot create new connection to %q", u.Redacted()), zap.Error(err))
	}
	c.Pool = pool
	c.Config = config
	return c
}
