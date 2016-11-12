package config

// ConfigurationReader defines the interface that provides access to all configurations parameters required by the service.
type ConfigurationReader interface {
	// GetListeningPort returns the port the application should start listening on.
	GetListeningPort() (int, error)

	// GetRootDirectory returns the root directory where files to be served are located.
	GetRootDirectory() (string, error)
}
