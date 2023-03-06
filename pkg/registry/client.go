package registry

import (
	"context"
	"fmt"

	"github.com/duyledat197/go-gen-tools/config"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type Options struct {
	//* enable tls
	IsEnableTLS bool
}

type ConsulClient struct {
	Client    *api.Client
	Config    *api.Config
	TLSConfig api.TLSConfig

	Logger  *zap.Logger
	Address *config.ConnectionAddr

	Options *Options
}

func (c *ConsulClient) Connect(ctx context.Context) error {
	cfg := api.DefaultConfig()
	cfg.Address = c.Address.GetConnectionString()

	if c.Options != nil {
		options := c.Options
		if options.IsEnableTLS {
			cfg.TLSConfig = c.TLSConfig
		}
	}

	client, err := api.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("create consul client error: %w", err)
	}
	c.Client = client
	c.Config = cfg
	c.Client.Connect()
	return nil
}

func (c *ConsulClient) Stop(ctx context.Context) error {
	return nil
}

func (c *ConsulClient) GetURL(serviceName string) string {
	return fmt.Sprintf("%s://%s/%s", "consul", c.Address, serviceName)
}
