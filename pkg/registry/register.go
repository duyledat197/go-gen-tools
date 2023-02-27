package registry

import (
	"fmt"
	"strconv"
	"time"

	"github.com/duyledat197/go-gen-tools/config"
	"github.com/duyledat197/go-gen-tools/utils"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

// ConsulRegister ...
type ConsulRegister struct {
	ServiceName                    string   // service name
	Tags                           []string // consul tags
	ConsulAddress                  *config.ConnectionAddr
	ServiceAddress                 *config.ConnectionAddr
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
	Client                         *api.Client
	Address                        string
	ID                             string
	Logger                         *zap.Logger
}

// Register ...
func (r *ConsulRegister) Register() error {
	ip := utils.GetLocalIP()
	port, _ := strconv.Atoi(r.ServiceAddress.Port)
	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", r.ServiceName, ip, port),
		Name:    r.ServiceName,
		Tags:    r.Tags,
		Port:    port,
		Address: ip,
		Check: &api.AgentServiceCheck{
			Interval:                       r.Interval.String(),
			GRPC:                           fmt.Sprintf("%v:%v/%v", ip, port, r.ServiceName),
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),
		},
	}
	r.Logger.Info("client connect consul with ID:", zap.String("registerID", reg.ID))
	r.ID = reg.ID
	return r.Client.Agent().ServiceRegister(reg)
}

func (r *ConsulRegister) Deregister() error {
	return r.Client.Agent().ServiceDeregister(r.ID)
}
