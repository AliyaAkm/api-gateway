package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
)

type HTTPConfig struct {
	ReadTimeout     time.Duration `env:"READ_TIMEOUT" envDefault:"5s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT" envDefault:"15s"`
	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT" envDefault:"60s"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"10s"`
}

type CORSConfig struct {
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" envSeparator:"," envDefault:"*"`
	AllowedMethods []string `env:"ALLOWED_METHODS" envSeparator:"," envDefault:"GET,POST,PUT,PATCH,DELETE,OPTIONS"`
	AllowedHeaders []string `env:"ALLOWED_HEADERS" envSeparator:"," envDefault:"Authorization,Content-Type,X-Request-ID"`
}

type UpstreamConfig struct {
	URL string `env:"URL"`
}

type Config struct {
	ServiceName string `env:"SERVICE_NAME" envDefault:"education-gateway"`
	HTTPAddr    string `env:"HTTP_ADDR" envDefault:":8090"`
	HTTP        HTTPConfig
	CORS        CORSConfig `envPrefix:"CORS_"`

	Auth         UpstreamConfig `envPrefix:"AUTH_SERVICE_"`
	Course       UpstreamConfig `envPrefix:"COURSE_SERVICE_"`
	Lesson       UpstreamConfig `envPrefix:"LESSON_SERVICE_"`
	Enrollment   UpstreamConfig `envPrefix:"ENROLLMENT_SERVICE_"`
	Progress     UpstreamConfig `envPrefix:"PROGRESS_SERVICE_"`
	Notification UpstreamConfig `envPrefix:"NOTIFICATION_SERVICE_"`
}

func ReadEnv() (*Config, error) {
	cfg := new(Config)

	opts := env.Options{
		RequiredIfNoDef: true,
	}

	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	required := map[string]string{
		"AUTH_SERVICE_URL":         c.Auth.URL,
		"COURSE_SERVICE_URL":       c.Course.URL,
		"LESSON_SERVICE_URL":       c.Lesson.URL,
		"ENROLLMENT_SERVICE_URL":   c.Enrollment.URL,
		"PROGRESS_SERVICE_URL":     c.Progress.URL,
		"NOTIFICATION_SERVICE_URL": c.Notification.URL,
	}

	for key, value := range required {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s is required", key)
		}
	}

	return nil
}
