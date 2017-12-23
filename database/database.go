package database

import (
	"fmt"

	"github.com/isaacjt/gatekeeper/config"
	"github.com/spf13/viper"
)

var (
	Db DatabaseBackend
)

func loadBackend(backendName string) (DatabaseBackend, error) {
	switch backendName {
	case "postgres":
		return &postgresBackend{}, nil
	case "sqlite":
		return &sqliteBackend{}, nil
	default:
		return nil, fmt.Errorf("invalid database backend: %s", backendName)
	}
}

func InitDatabase() error {
	var err error
	Db, err = loadBackend(viper.GetString(config.ConfigDatabaseBackend))
	if err != nil {
		return err
	}
	return Db.Init()
}
