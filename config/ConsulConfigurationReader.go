package config

import "github.com/microbusinesses/Micro-Businesses-Core/common/config"

// ConsulConfigurationReader implements ConfigurationReader using Consul to provide access to all configurations parameters required by the service.
type ConsulConfigurationReader struct {
	ConsulAddress           string
	ConsulScheme            string
	ListeningPortToOverride int
}

const serviceListeningPortKey = "services/uihost-service/endpoint/listening-port"

// GetListeningPort returns the port the service should listen on to serve the HTTP request
func (consul ConsulConfigurationReader) GetListeningPort() (int, error) {
	if consul.ListeningPortToOverride != 0 {
		return consul.ListeningPortToOverride, nil

	}
	consulHelper := config.ConsulHelper{ConsulAddress: consul.ConsulAddress, ConsulScheme: consul.ConsulScheme}

	return consulHelper.GetInt(serviceListeningPortKey)
}
