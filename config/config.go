package config

import "github.com/sebomancien/goth-template/internal/server"

type Config struct {
	Server server.Config
}

func Load() (*Config, error) {
	return &Config{
		Server: server.Config{
			Host: "localhost",
			Port: 3000,
		},
	}, nil
}
