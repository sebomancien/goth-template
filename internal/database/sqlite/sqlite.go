package sqlite

import (
	"database/sql"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
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
