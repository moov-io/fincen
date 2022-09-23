// generated-from:2a176987fda9186846b2dd87f406a45da445a8c8bf94d05b1571978d27f11019 DO NOT REMOVE, DO UPDATE

package service_test

import (
	"testing"

	"github.com/moov-io/fincen/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvironment(t *testing.T) {
	env, err := service.NewEnvironment(nil)

	require.NoError(t, err)
	t.Cleanup(env.Shutdown)
	require.NotNil(t, env.Logger)
	require.NotNil(t, env.Config)
	require.NotNil(t, env.InternalClient)
	require.NotNil(t, env.TimeService)
	require.NotNil(t, env.PublicRouter)

	t.Run("servers - run and shut down", func(t *testing.T) {
		listener := service.NewTerminationListener()
		shutdown := env.RunServers(listener)
		shutdown()

		msg := <-listener
		require.NotNil(t, msg)
		assert.Contains(t, msg.Error(), "http: Server closed")
	})
}
