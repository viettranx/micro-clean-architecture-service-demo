package cmd

import (
	"demo-service/common"
	"flag"
	sctx "github.com/viettranx/service-context"
)

type config struct {
	grpcPort          int    // for server port listening
	grpcServerAddress string // for client make grpc client connection
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ID() string {
	return common.KeyCompConf
}

func (c *config) InitFlags() {
	flag.IntVar(
		&c.grpcPort,
		"grpc-port",
		3300,
		"gRPC Port. Default: 3300",
	)

	flag.StringVar(
		&c.grpcServerAddress,
		"grpc-server-address",
		"localhost:3300",
		"gRPC server address. Default: localhost:3300",
	)
}

func (c *config) Activate(_ sctx.ServiceContext) error {
	return nil
}

func (c *config) Stop() error {
	return nil
}

func (c *config) GetGRPCPort() int {
	return c.grpcPort
}

func (c *config) GetGRPCServerAddress() string {
	return c.grpcServerAddress
}
