package database

import (
	"fmt"

	"github.com/spf13/viper"
	"is.having.coffee/gatekeeper/gatekeeper-server/config"
)

var (
	Db DatabaseBackend
)

func loadBackend(backendName string) (DatabaseBackend, error) {
	switch backendName {
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
