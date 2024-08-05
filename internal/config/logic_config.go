package config

func (c *ConfigService) GetDuration() int {
	return c.config.TickDuration
}
