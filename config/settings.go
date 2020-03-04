package config

type Settings struct {
	Environment string `json:"ENVIRONMENT" env:"ENVIRONMENT"`
	LoggingLevel string `json:"LOGGING_LEVEL" env:"LOGGING_LEVEL"`

	HttpHost string `json:"HTTP_HOST" env:"HTTP_HOST"`
	HttpPort string `json:"HTTP_PORT" env:"HTTP_PORT"`

	DbType     string `json:"DB_TYPE" env:"DB_TYPE"`
	DbHost     string `json:"DB_HOST" env:"DB_HOST"`
	DbDataBase string `json:"DB_DATABASE" env:"DB_DATABASE"`
	DbUsername string `json:"DB_USERNAME" env:"DB_USERNAME"`
	DbPassword string `json:"DB_PASSWORD" env:"DB_PASSWORD"`
	DbPort     string `json:"DB_PORT" env:"DB_PORT"`
	DbSchema   string `json:"DB_SCHEMA" env:"DB_SCHEMA"`
	DbSSLMode  string `json:"DB_SSL_MODE" env:"DB_SSL_MODE"`
}
