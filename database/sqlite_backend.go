package database

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"is.having.coffee/gatekeeper/gatekeeper-server/config"
)

type SqliteBackend interface {
	DatabaseBackend
}

type sqliteBackend struct {
	SqliteBackend
	db *sql.DB
}

func (backend *sqliteBackend) Init() error {
	dsn := viper.GetString(config.ConfigDatabaseConnString)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}
	backend.db = db
	initScript, err := ioutil.ReadFile("resource/sql/init_sqlite3.sql")
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return nil
	}
	if _, err := tx.Exec(string(initScript)); err != nil {
		return err
	}
	return tx.Commit()
}

func (backend sqliteBackend) Close() error {
	if backend.db != nil {
		return backend.db.Close()
	}
	return nil
}

func (backend sqliteBackend) GetDb() *sql.DB {
	return backend.db
}
