// generated-from:d1859aa7791ddd53eff7074f9c0c47fee8df416501547d4c6e8b713c60e74a4d DO NOT REMOVE, DO UPDATE

package service_test

import (
	"fmt"
	"testing"

	"github.com/moov-io/base/config"
	"github.com/moov-io/base/log"
	"github.com/moov-io/fincen/pkg/service"
	"github.com/stretchr/testify/require"
)

func Test_ConfigLoading(t *testing.T) {
	logger := log.NewNopLogger()

	ConfigService := config.NewService(logger)

	gc := &service.GlobalConfig{}
	err := ConfigService.Load(gc)
	fmt.Println(err)
	require.Nil(t, err)
}
