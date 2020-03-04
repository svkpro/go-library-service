package config

type Settings struct {
	Environment string `json:"ENVIRONMENT" env:"ENVIRONMENT"`
	LoggingLevel string `json:"LOGGING_LEVEL" env:"LOGGING_LEVEL"`

	HttpHost string `json:"HTTP_HOST" env:"HTTP_HOST"`
	HttpPort string `json:"HTTP_PORT" env:"HTTP_PORT"`
}
