// generated-from:950a5a6611312ac571326fead04ceaeb6428e3570b8243d80e2a9ea456c341d2 DO NOT REMOVE, DO UPDATE

package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/moov-io/base/admin"
	"github.com/moov-io/base/log"

	_ "github.com/moov-io/fincen"
)

// RunServers - Boots up all the servers and awaits till they are stopped.
func (env *Environment) RunServers(terminationListener chan error) func() {

	adminServer := bootAdminServer(terminationListener, env.Logger, env.Config.Servers.Admin)
	handler := env.PublicRouter

	shutdownPublicServer := func() {}
	if env.Config.Servers.Public != nil {
		_, shutdownPublicServer = bootHTTPServer("public", handler, terminationListener, env.Logger, *env.Config.Servers.Public)
	}

	return func() {
		adminServer.Shutdown()
		shutdownPublicServer()
	}
}

// Deprecated and will go away once TLS is out to everyone
func bootHTTPServer(name string, routes http.Handler, errs chan<- error, logger log.Logger, config HTTPConfig) (*http.Server, func()) {

	// Create main HTTP server
	serve := &http.Server{
		Addr:    config.Bind.Address,
		Handler: routes,
		TLSConfig: &tls.Config{
			InsecureSkipVerify:       false,
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
		},
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Start main HTTP server
	go func() {
		logger.Info().Log(fmt.Sprintf("%s listening on %s for HTTP", name, config.Bind.Address))
		if err := serve.ListenAndServe(); err != nil {
			errs <- logger.Fatal().LogErrorf("problem starting http: %w", err).Err()
		}
	}()

	shutdownServer := func() {
		if err := serve.Shutdown(context.Background()); err != nil {
			logger.Fatal().LogErrorf("shutting down: %v", err)
		}
	}

	return serve, shutdownServer
}

func bootAdminServer(errs chan<- error, logger log.Logger, config HTTPConfig) *admin.Server {
	adminServer, err := admin.New(admin.Opts{
		Addr: config.Bind.Address,
	})
	if err != nil {
		errs <- logger.Fatal().LogErrorf("problem creating admin server: %v", err).Err()
	}

	go func() {
		logger.Info().Log(fmt.Sprintf("listening on %s", adminServer.BindAddr()))
		if err := adminServer.Listen(); err != nil {
			errs <- logger.Fatal().LogErrorf("problem starting admin http: %w", err).Err()
		}
	}()

	return adminServer
}
