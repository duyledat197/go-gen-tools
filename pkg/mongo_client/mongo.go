package mongo_client

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/duyledat197/go-gen-tools/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type Options struct {
	//* enable tls
	IsEnableTLS bool

	//* enable load balancer
	IsEnableLoadBalancer bool

	//* enable replicaset
	IsEnableRelicaSet bool
}

type MongoClient struct {
	Database      *config.Database
	Client        *mongo.Client
	clientOptions *options.ClientOptions

	Logger     *zap.Logger
	TLSConfig  tls.Config
	ReplicaSet string

	Options *Options
}

func (c *MongoClient) Connect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, c.Database.Timeout)
	defer cancel()
	clientOpts := options.Client()

	//* set up mongo client options
	clientOpts.SetConnectTimeout(c.Database.Timeout)
	clientOpts.SetMaxConnIdleTime(15 * time.Second)
	clientOpts.SetMaxConnecting(uint64(c.Database.MaxConnection))
	clientOpts.SetRetryReads(true)
	clientOpts.SetRetryWrites(true)

	if c.Options != nil {
		options := c.Options
		if options.IsEnableTLS {
			clientOpts.SetTLSConfig(&c.TLSConfig)
		}

		clientOpts.SetLoadBalanced(options.IsEnableLoadBalancer)

		if options.IsEnableRelicaSet {
			clientOpts.SetReplicaSet(c.ReplicaSet)
		}
	}

	if err := clientOpts.Validate(); err != nil {
		return fmt.Errorf("connect mongo error: validate : %w", err)
	}
	client, err := mongo.Connect(ctx, clientOpts.ApplyURI(c.Database.GetConnectionString()))
	if err != nil {
		return fmt.Errorf("connect mongo error: connect: %w", err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("connect mongo error: ping: %w", err)
	}
	c.Client = client
	c.clientOptions = clientOpts
	return nil
}

func (c *MongoClient) Stop(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}
