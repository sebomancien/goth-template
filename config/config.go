package config

import "github.com/sebomancien/goth-template/internal/db"

type Config struct {
	Database db.Config
}

func Load() (*Config, error) {
	return &Config{
		Database: db.Config{
			Type: db.Sqlite,
			Name: "/data/database.sqlite",
		},
	}, nil
}
