package config

import "github.com/micro-business/Micro-Business-Core/common/config"

// ConsulConfigurationReader implements ConfigurationReader using Consul to provide access to all configurations parameters required by the service.
type ConsulConfigurationReader struct {
	ConsulAddress           string
	ConsulScheme            string
	ListeningPortToOverride int
	RootDirectoryToOverride string
}

const serviceListeningPortKey = "services/uihost-service/endpoint/listening-port"
const rootDirectoryKey = "services/uihost-service/endpoint/root-directory"

// GetListeningPort returns the port the service should listen on to serve the HTTP request
func (consul ConsulConfigurationReader) GetListeningPort() (int, error) {
	if consul.ListeningPortToOverride != 0 {
		return consul.ListeningPortToOverride, nil

	}

	consulHelper := config.ConsulHelper{ConsulAddress: consul.ConsulAddress, ConsulScheme: consul.ConsulScheme}

	return consulHelper.GetInt(serviceListeningPortKey)
}

// GetRootDirectory returns the root directory where files to be served are located.
func (consul ConsulConfigurationReader) GetRootDirectory() (string, error) {
	if len(consul.RootDirectoryToOverride) != 0 {
		return consul.RootDirectoryToOverride, nil
	}

	consulHelper := config.ConsulHelper{ConsulAddress: consul.ConsulAddress, ConsulScheme: consul.ConsulScheme}

	return consulHelper.GetString(rootDirectoryKey)
}
