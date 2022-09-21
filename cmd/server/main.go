// generated-from:f953aec60bb5d48612763d5ba805f8f394b4cbf292e11dd1942281ee9c8b166c DO NOT REMOVE, DO UPDATE

package main

import (
	"os"

	"github.com/moov-io/base/log"

	"github.com/moov-io/fincen"
	"github.com/moov-io/fincen/pkg/service"
)

func main() {
	logger := log.NewDefaultLogger().
		Set("app", log.String("server")).
		Set("version", log.String(fincen.Version))

	env := &service.Environment{
		Logger: logger,
	}

	_, err := service.NewEnvironment(env)
	if err != nil {
		logger.Fatal().LogErrorf("Error loading up environment: %v", err)
		os.Exit(1)
	}

	termListener := service.NewTerminationListener()

	stopServers := env.RunServers(termListener)

	service.AwaitTermination(env.Logger, termListener)

	stopServers()
	env.Shutdown()
}
