// Copyright 2022 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

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
