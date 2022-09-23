// generated-from:728e0fccc6600a6a5cfb0e340a5432625730ffb2e05575e0cfe522989df22164 DO NOT REMOVE, DO UPDATE

package service

type GlobalConfig struct {
	Fincen Config
}

// Config defines all the configuration for the app
type Config struct {
	Servers ServerConfig
	Clients *ClientConfig
}

// ServerConfig - Groups all the http configs for the servers and ports that get opened.
type ServerConfig struct {
	Public *HTTPConfig
	Admin  HTTPConfig
}

// HTTPConfig configuration for running an http server
type HTTPConfig struct {
	Bind BindAddress
}

// BindAddress specifies where the http server should bind to.
type BindAddress struct {
	Address string
}
