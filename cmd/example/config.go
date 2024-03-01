package main

import (
	"github.com/Snork2017/example-srv/cmd/example/handlers"
	"github.com/Snork2017/example-srv/internal/pkg/storage/redis"
	"github.com/caarlos0/env/v10"
)

type Config struct {
	// import examplestore.Config
	// some usable configs for the example service.
	Port string

	GRPCConf handlers.Config
	Redis    redis.Config
}

func (c *Config) Load() error {
	return env.Parse(c)
}
