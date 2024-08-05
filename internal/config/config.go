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

// TODO понять какая разница между вызовом конструктора и вызовом хука fx.Hook (тк вообще не определен порядок)
func NewConfigService(lc fx.Lifecycle) (*ConfigService, error) {
	c := &ConfigService{}
	slog.Info("NewConfigService constructor")

	configPath, exists := os.LookupEnv(urlCheckerConfig)
	if !exists {
		configPath = "urlCheckerConfig.json"
	}

	conf := commonConfig{}

	b, err := os.ReadFile(configPath)
	if err != nil {
		return c, errors.Wrap(err, "os.ReadFile")
	}

	err = json.Unmarshal(b, &conf)
	if err != nil {
		return c, errors.Wrap(err, "Unmarshal"+" config path "+urlCheckerConfig)
	}

	c.config = conf

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				slog.Info("NewConfigService OnStart")

				return nil
			},
			OnStop: func(ctx context.Context) error {
				slog.Info("NewConfigService OnStop")
				return nil
			},
		},
	)

	return c, nil
}
