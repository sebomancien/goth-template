package config

import (
	"github.com/sebomancien/goth-template/internal/database"
	"github.com/sebomancien/goth-template/internal/server"
)

type Config struct {
	Database database.Config
	Server   server.Config
}

func Load() (*Config, error) {
	return &Config{
		Database: database.Config{
			Type: database.Sqlite,
			Name: "/data/database.sqlite",
		},
		Server: server.Config{
			Host: "localhost",
			Port: 3000,
		},
	}, nil
}
