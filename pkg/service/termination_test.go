// generated-from:82de90d6a7f71cb4a4b5e2e5f01606226cfa1e6f5567655b25e0a9f4df71cff8 DO NOT REMOVE, DO UPDATE

package service_test

import (
	"errors"
	"testing"

	"github.com/moov-io/base/log"
	"github.com/moov-io/fincen/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTermination(t *testing.T) {
	listener := service.NewTerminationListener()
	err := make(chan error)
	go func() {
		err <- service.AwaitTermination(log.NewNopLogger(), listener)
	}()
	listener <- errors.New("foo")

	got := <-err
	require.Error(t, got)
	assert.Contains(t, got.Error(), "Terminated: foo")
}
