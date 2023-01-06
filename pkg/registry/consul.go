package registry

import (
	"fmt"
	"time"

	"github.com/duyledat197/go-gen-tools/utils"

	"github.com/hashicorp/consul/api"
)

// NewClient returns a new Client with connection to consul
func NewClient(addr string) (*api.Client, error) {
	cfg := api.DefaultConfig()
	cfg.Address = addr

	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return c, nil
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
}

// NewConsulRegister ...
func NewConsulRegister(consulClient *api.Client, serviceName string, servicePort int, tags []string) *ConsulRegister {
	return &ConsulRegister{
		ServiceName:                    serviceName,
		Tags:                           tags,
		ServicePort:                    servicePort,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
		Client:                         consulClient,
	}
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
	fmt.Println("client connect consul with ID:", reg.ID)
	return reg.ID, r.Client.Agent().ServiceRegister(reg)
}

func (r *ConsulRegister) Deregister(id string) error {
	return r.Client.Agent().ServiceDeregister(id)
}
