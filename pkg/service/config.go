// generated-from:55d46bf0f65f2d352c704e686f8bcfb1bffd75f5089be80f8720a046a61b1d79 DO NOT REMOVE, DO UPDATE

package service

import (
	"github.com/moov-io/base/config"
	"github.com/moov-io/base/log"

	_ "github.com/moov-io/fincen"
)

func LoadConfig(logger log.Logger) (*Config, error) {
	configService := config.NewService(logger)

	global := &GlobalConfig{}
	if err := configService.Load(global); err != nil {
		return nil, err
	}

	cfg := &global.Fincen

	return cfg, nil
}
