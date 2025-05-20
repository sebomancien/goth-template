package db

import (
	"fmt"
	"io/fs"

	"github.com/sebomancien/goth-template/internal/db/sqlite"
	"github.com/sebomancien/goth-template/internal/model"
)

type Type int

const (
	Sqlite Type = iota
)

type Config struct {
	Type Type
	Name string
}

type Database interface {
	Close()
	AddLog(level model.LogLevel, message string) error
	GetLogs() ([]model.Log, error)
	GetTables() ([]string, error)
}

func Open(config *Config, migration fs.FS) (Database, error) {
	switch config.Type {
	case Sqlite:
		return sqlite.Open(config.Name, migration)
	default:
		return nil, fmt.Errorf("unsupported database type: %v", config.Type)
	}
}
