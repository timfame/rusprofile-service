package config

import "github.com/timfame/rusprofile-service/pkg/env"

const (
	GRPCPortEnv        = "GRPC_PORT"
	GRPCGatewayPortEnv = "GRPC_GATEWAY_PORT"
)

type Grpc struct {
	Port        string `json:"port"`
	GatewayPort string `json:"gateway_port"`
}

func (s *Grpc) Init() (err error) {
	s.Port, err = env.GetString(GRPCPortEnv)
	if err != nil {
		return
	}
	s.GatewayPort, err = env.GetString(GRPCGatewayPortEnv)
	return
}
