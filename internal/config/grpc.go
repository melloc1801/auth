package config

import (
	"errors"
	"net"
	"os"
)

const (
	grpcPortName = "GRPC_PORT"
	grpcHostName = "GRPC_HOST"
)

type GRPCConfig interface {
	Address() string
}
type grpcConfig struct {
	port string
	host string
}

func NewGrpcConfig() (GRPCConfig, error) {
	port := os.Getenv(grpcPortName)
	if len(port) == 0 {
		return nil, errors.New("GRPC_PORT not found")
	}

	host := os.Getenv(grpcHostName)
	if len(host) == 0 {
		return nil, errors.New("GRPC_HOST not found")
	}

	return &grpcConfig{port: port, host: host}, nil
}

func (g *grpcConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}
