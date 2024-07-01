package config

func (c *ConfigService) MainServiceConfig() (string, int) {
	return c.config.MainServiceConfig.Host, c.config.MainServiceConfig.Port
}

func (c *ConfigService) PingServiceConfig() (string, int) {
	return c.config.PingServiceConfig.Host, c.config.PingServiceConfig.Port
}
