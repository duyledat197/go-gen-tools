package redis_client

import (
	"context"
	"fmt"
	"time"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type RedisClient struct {
	Client   *redis.Client
	Database *config.Database

	Logger *zap.Logger
}

func (c *RedisClient) Connect(ctx context.Context) error {
	options := &redis.Options{
		Addr:            c.Database.GetConnectionString(),
		Username:        c.Database.UserName,
		Password:        c.Database.Password,
		MaxRetries:      5,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 512 * time.Millisecond,
		DialTimeout:     5 * time.Second,
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    5 * time.Second,
		PoolFIFO:        false,
		PoolSize:        c.Database.MaxConnection,
	}
	client := redis.NewClient(options)
	if cmd := client.Ping(ctx); cmd.Err() != nil {
		return fmt.Errorf("connect redis error: %w", cmd.Err())
	}
	c.Client = client
	return nil
}

func (c *RedisClient) Stop(ctx context.Context) error {
	return c.Client.Close()
}

type Options struct {
	TTL                    *time.Time
	Sort                   *redis.Sort
	GeoSearchLocationQuery *redis.GeoSearchLocationQuery
	GeoRadiusQuery         *redis.GeoRadiusQuery
	GeoSearchQuery         *redis.GeoSearchQuery
	GeoSearchStoreQuery    *redis.GeoSearchStoreQuery
}

func GetOption(opts []*Options) *Options {
	if len(opts) == 0 || opts[0] == nil {
		return nil
	}
	return opts[0]
}

func (c *RedisClient) Get(ctx context.Context, key string, opts ...*Options) (protoreflect.ProtoMessage, error) {
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	var result protoreflect.ProtoMessage
	if err := ConvertToProtoMessage(val, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *RedisClient) Set(ctx context.Context, key string, val protoreflect.ProtoMessage, opts ...*Options) error {
	opt := GetOption(opts)
	data, err := ConvertToString(val)
	if err != nil {
		return err
	}
	if opt != nil && opt.TTL != nil {
		if _, err := c.Client.SetNX(ctx, key, data, time.Duration(opt.TTL.Nanosecond())).Result(); err != nil {
			return err
		}
	}

	if _, err := c.Client.SetNX(ctx, key, data, redis.KeepTTL).Result(); err != nil {
		return err
	}

	return nil
}

func ConvertToString[T protoreflect.ProtoMessage](val T) (string, error) {
	b, err := proto.Marshal(val)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ConvertToProtoMessage[T protoreflect.ProtoMessage](data string, val T) error {
	err := proto.Unmarshal([]byte(data), val)
	if err != nil {
		return err
	}
	return nil
}
