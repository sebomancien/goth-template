package models

import "time"

type LogLevel int

const (
	LogLevelVerbose LogLevel = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
	LogLevelFatal
)

type Log struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
}
