package sqlite

import (
	"database/sql"
	"io/fs"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sebomancien/goth-template/internal/model"
)

type Database struct {
	db *sql.DB
}

func Open(file string, migration fs.FS) (*Database, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	err = runMigration(db, migration)
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func runMigration(db *sql.DB, migration fs.FS) error {
	source, err := iofs.New(migration, "migration/sqlite")
	if err != nil {
		return err
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", source, "sqlite3", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) AddLog(level model.LogLevel, message string) error {
	_, err := d.db.Exec("INSERT INTO logs (timestamp, level, message) VALUES (?, ?, ?)", time.Now(), level, message)
	return err
}

func (d *Database) GetLogs() ([]model.Log, error) {
	rows, err := d.db.Query(`SELECT timestamp, level, message FROM logs`)
	if err != nil {
		return nil, err
	}

	var logs []model.Log
	for rows.Next() {
		var log model.Log
		err := rows.Scan(&log.Timestamp, &log.Level, &log.Message)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (d *Database) GetTables() ([]string, error) {
	rows, err := d.db.Query(`SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tables, nil
}
