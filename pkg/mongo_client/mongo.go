package mongo_client

import (
	"context"
	"time"

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
	Client        *mongo.Client
	ClientOptions *options.ClientOptions

	Timeout       time.Duration
	MaxConnection uint64

	ConnectionURI string
	Logger        *zap.Logger
	Creds         options.Credential
	ReplicaSet    string

	Options *Options
}

func (c *MongoClient) Init(ctx context.Context) *MongoClient {
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	clientOpts := options.Client()

	//* set up mongo client options
	clientOpts.SetConnectTimeout(c.Timeout)
	clientOpts.SetMaxConnIdleTime(15 * time.Second)
	clientOpts.SetMaxConnecting(c.MaxConnection)
	clientOpts.SetRetryReads(true)
	clientOpts.SetRetryWrites(true)

	if c.Options != nil {
		options := c.Options
		if options.IsEnableTLS {
			clientOpts.SetAuth(c.Creds)
		}

		clientOpts.SetLoadBalanced(options.IsEnableLoadBalancer)

		if options.IsEnableRelicaSet {
			clientOpts.SetReplicaSet(c.ReplicaSet)
		}
	}

	if err := clientOpts.Validate(); err != nil {
		c.Logger.Panic("connect mongo error: validate ", zap.Error(err))
	}
	client, err := mongo.Connect(ctx, clientOpts.ApplyURI(c.ConnectionURI))
	if err != nil {
		c.Logger.Panic("connect mongo error: connect ", zap.Error(err))
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		c.Logger.Panic("connect mongo error: ping ", zap.Error(err))
	}
	c.Client = client
	c.ClientOptions = clientOpts
	return c
}
