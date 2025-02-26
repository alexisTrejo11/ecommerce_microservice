package config

import (
	"os"
)

type LoggerConfig struct {
	Environment  string
	LogLevel     string
	IsProduction bool
	LogstashURL  string
	UseELK       bool
}

func NewLoggerConfig() LoggerConfig {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	logstashURL := os.Getenv("LOGSTASH_URL")
	if logstashURL == "" {
		logstashURL = "logstash:5000"
	}

	useELK := os.Getenv("USE_ELK") == "true"

	return LoggerConfig{
		Environment:  env,
		LogLevel:     logLevel,
		IsProduction: env == "production",
		LogstashURL:  logstashURL,
		UseELK:       useELK,
	}
}
