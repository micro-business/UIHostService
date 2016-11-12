package config

// ConfigurationReader defines the interface that provides access to all configurations parameters required by the service.
type ConfigurationReader interface {
	// GetListeningPort returns the port the application should start listening on.
	GetListeningPort() (int, error)
}
