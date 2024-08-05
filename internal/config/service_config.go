package config

type serviceConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type commonConfig struct {
	MainServiceConfig serviceConfig `json:"mainServiceConfig"`
	PingServiceConfig serviceConfig `json:"pingServiceConfig"`
	TickDuration      int           `json:"tickDuration"`
}
