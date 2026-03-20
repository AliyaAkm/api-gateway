package config

import (
	"fmt"
	"net"
	"net/url"
	"os"
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
	URL string `env:"URL" envDefault:""`
}

func (u UpstreamConfig) BaseURL() string {
	value := strings.TrimSpace(u.URL)
	if value == "" {
		return ""
	}

	parsed, err := url.Parse(value)
	if err == nil && parsed.Scheme != "" && parsed.Host != "" {
		return strings.TrimRight(value, "/")
	}

	return "http://" + strings.TrimRight(value, "/")
}

type Config struct {
	ServiceName string `env:"SERVICE_NAME" envDefault:"education-gateway"`
	HTTPAddr    string `env:"HTTP_ADDR" envDefault:":8090"`
	HTTP        HTTPConfig
	CORS        CORSConfig `envPrefix:"CORS_"`

	Auth         UpstreamConfig `envPrefix:"AUTH_SERVICE_"`
	Curriculum   UpstreamConfig `envPrefix:"CURRICULUM_SERVICE_"`
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
	configured := []string{
		c.Auth.BaseURL(),
		c.Curriculum.BaseURL(),
		c.Lesson.BaseURL(),
		c.Enrollment.BaseURL(),
		c.Progress.BaseURL(),
		c.Notification.BaseURL(),
	}

	for _, value := range configured {
		if strings.TrimSpace(value) != "" {
			return nil
		}
	}

	return fmt.Errorf("at least one upstream service URL is required")
}

func (c Config) ListenAddr() string {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		return c.HTTPAddr
	}

	return net.JoinHostPort("", port)
}
