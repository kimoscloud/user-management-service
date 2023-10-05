package configuration

import (
	"os"
	"strconv"
)

type HttpServerConfig struct {
	port uint
}

// Getters for HttpServerConfig

func (httpServerConfig *HttpServerConfig) GetPort() uint {
	return httpServerConfig.port
}

var httpServerConfiguration *HttpServerConfig

func GetHttpServerConfig() *HttpServerConfig {
	if httpServerConfiguration == nil {
		initHttpServerConfig()
	}
	return httpServerConfiguration
}

func initHttpServerConfig() {
	portStr := os.Getenv("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 10101
	}

	httpServerConfiguration = &HttpServerConfig{
		port: uint(port),
	}
}
