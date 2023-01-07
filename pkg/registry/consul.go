package registry

import (
	"fmt"
	"time"

	"github.com/duyledat197/go-gen-tools/utils"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type ConsulClient struct {
	Client *api.Client
	Config *api.Config

	Logger  *zap.Logger
	Address string
}

func (c *ConsulClient) Init() *ConsulClient {
	cfg := api.DefaultConfig()
	cfg.Address = c.Address

	client, err := api.NewClient(cfg)
	if err != nil {
		c.Logger.Panic("create consul client error: ", zap.Error(err))
	}
	c.Client = client
	c.Config = cfg

	return c
}

func (c *ConsulClient) GetURL(serviceName string) string {
	return fmt.Sprintf("%s://%s/%s", "consul", c.Address, serviceName)
}

// ConsulRegister ...
type ConsulRegister struct {
	ServiceName                    string   // service name
	Tags                           []string // consul tags
	ServicePort                    int      //service port
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
	Client                         *api.Client
	Address                        string

	Logger *zap.Logger
}

// Register ...
func (r *ConsulRegister) Register() (string, error) {
	IP := utils.GetLocalIP()
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", r.ServiceName, IP, r.ServicePort),
		Name:    r.ServiceName,
		Tags:    r.Tags,
		Port:    r.ServicePort,
		Address: IP,
		Check: &api.AgentServiceCheck{
			Interval:                       r.Interval.String(),
			GRPC:                           fmt.Sprintf("%v:%v/%v", IP, r.ServicePort, r.ServiceName),
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),
		},
	}
	r.Logger.Info("client connect consul with ID:", zap.String("registerID", reg.ID))
	return reg.ID, r.Client.Agent().ServiceRegister(reg)
}

func (r *ConsulRegister) Deregister(id string) error {
	return r.Client.Agent().ServiceDeregister(id)
}
