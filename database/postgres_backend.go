package database

import (
	"database/sql"
	"io/ioutil"
	"log"

	"github.com/isaacjt/gatekeeper/config"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type PostgresBackend interface {
	DatabaseBackend
}

type postgresBackend struct {
	PostgresBackend
	db *sql.DB
}

func (backend *postgresBackend) Init() error {
	dsn := viper.GetString(config.ConfigDatabaseConnString)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	backend.db = db
	initScript, err := ioutil.ReadFile("resource/sql/init_psql.sql")
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

func (backend postgresBackend) Close() error {
	if backend.db != nil {
		return backend.db.Close()
	}
	return nil
}

func (backend postgresBackend) GetDb() *sql.DB {
	return backend.db
}
