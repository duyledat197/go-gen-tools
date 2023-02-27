package postgres_client

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/duyledat197/go-gen-tools/config"
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
	Pool     *pgxpool.Pool
	Database *config.Database
	config   *pgxpool.Config

	Logger  *zap.Logger
	Options *Options
}

func (c *PostgresClient) Connect(ctx context.Context) error {
	connectionString := c.Database.GetConnectionString()
	u, err := url.Parse(connectionString)
	if err != nil {
		return fmt.Errorf("cannot create new connection to Postgres (failed to parse URI)", err)
	}
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return fmt.Errorf("cannot read PG_CONNECTION_URI", err)
	}

	config.MaxConns = int32(c.Database.MaxConnection)
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

	c.Logger.Info("NewConnectionPool max connection", zap.Int("max connection", c.Database.MaxConnection))

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("cannot create new connection to %q", u.Redacted()), err)
	}
	c.Pool = pool
	c.config = config
	return nil
}

func (c *PostgresClient) Stop(ctx context.Context) error {
	c.Pool.Close()
	return nil
}
