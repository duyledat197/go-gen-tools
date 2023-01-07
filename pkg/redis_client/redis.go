package redis_client

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type RedisClient struct {
	Client  *redis.Client
	Options *redis.Options
}

func NewRedisClient(options *redis.Options) *RedisClient {
	ctx := context.Background()
	client := redis.NewClient(options)
	if cmd := client.Ping(ctx); cmd.Err() != nil {
		panic(cmd.Err())
	}
	return &RedisClient{
		Client:  client,
		Options: options,
	}
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
