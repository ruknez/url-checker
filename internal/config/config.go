package config

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/pkg/errors"
	"go.uber.org/fx"
)

const urlCheckerConfig = "URL_CHECKER_CONFIG"

type ConfigService struct {
	config commonConfig
}

func NewConfigService(lc fx.Lifecycle) *ConfigService {
	c := &ConfigService{}
	slog.Info("NewConfigService constructor")

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				slog.Info("NewConfigService OnStart")
				configPath, exists := os.LookupEnv(urlCheckerConfig)
				if !exists {
					configPath = "urlCheckerConfig.json"
				}

				conf := commonConfig{}

				b, err := os.ReadFile(configPath)
				if err != nil {
					return errors.Wrap(err, "os.ReadFile")
				}

				err = json.Unmarshal(b, &conf)
				if err != nil {
					return errors.Wrap(err, "Unmarshal"+" config path "+urlCheckerConfig)
				}

				c.config = conf

				return nil
			},
			OnStop: func(ctx context.Context) error {
				slog.Info("NewConfigService OnStop")
				return nil
			},
		},
	)

	return c
}
