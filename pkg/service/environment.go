// generated-from:859eb7e73edee443e3c17b3f44ad9fa410e70a0b22917ec40a21d123924d79e4 DO NOT REMOVE, DO UPDATE

package service

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moov-io/base/log"
	"github.com/moov-io/base/stime"

	_ "github.com/moov-io/fincen"
)

// Environment - Contains everything thats been instantiated for this service.
type Environment struct {
	Logger         log.Logger
	Config         *Config
	TimeService    stime.TimeService
	DB             *sql.DB
	InternalClient *http.Client

	PublicRouter *mux.Router
	Shutdown     func()
}

// NewEnvironment - Generates a new default environment. Overrides can be specified via configs.
func NewEnvironment(env *Environment) (*Environment, error) {
	if env == nil {
		env = &Environment{}
	}

	env.Shutdown = func() {}

	if env.Logger == nil {
		env.Logger = log.NewDefaultLogger()
	}

	if env.Config == nil {
		cfg, err := LoadConfig(env.Logger)
		if err != nil {
			return nil, err
		}

		env.Config = cfg
	}

	if env.TimeService == nil {
		env.TimeService = stime.NewSystemTimeService()
	}

	if env.InternalClient == nil {
		env.InternalClient = NewInternalClient(env.Logger, env.Config.Clients, "internal-client")
	}

	// router
	if env.PublicRouter == nil {
		env.PublicRouter = mux.NewRouter()
	}

	// configure custom handlers
	ConfigureHandlers(env.PublicRouter)

	return env, nil
}
