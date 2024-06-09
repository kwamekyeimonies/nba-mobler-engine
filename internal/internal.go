package internal

import (
	"errors"
	"github.com/kwamekyeimonies/nba-mobler-engine/config"
)

func InitializeDependencies() error {
	if config.WithConfigInjector = config.NewConfigurationInjector(); config.WithConfigInjector == nil {
		return errors.New("failed to initialize config")
	}

	return nil
}
