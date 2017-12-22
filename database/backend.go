package database

import "database/sql"

type DatabaseBackend interface {
	Init() error
	Close() error
	GetDb() *sql.DB
}
