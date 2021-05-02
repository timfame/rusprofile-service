package logger

import (
	"go.uber.org/zap"
)

const (
	DEBUG = iota
	RELEASE
)

type Config struct {
	serviceName string
	zapConfig   *zap.Config
	level       int
}

type Option func(*Config)

func WithServiceName(name string) Option {
	return func(c *Config) {
		c.serviceName = name
	}
}

func WithZapConfig(zapCfg *zap.Config) Option {
	return func(c *Config) {
		c.zapConfig = zapCfg
	}
}

func WithDebugLevel() Option {
	return func(c *Config) {
		c.level = DEBUG
	}
}

func WithReleaseLevel() Option {
	return func(c *Config) {
		c.level = RELEASE
	}
}
