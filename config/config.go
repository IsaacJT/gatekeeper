package config

import (
	"github.com/spf13/viper"
)

const (
	ConfigDatabaseBackend    = "db_backend"
	ConfigDatabaseConnString = "db_connection_string"
	ConfigConsole            = "console"
	ConfigListenSocket       = "listen"
)

var (
	defaults = map[string]interface{}{
		ConfigDatabaseBackend:    "sqlite",
		ConfigDatabaseConnString: "gatekeeper.sqlite3",
		ConfigConsole:            true,
		ConfigListenSocket:       "localhost:8080",
	}
)

func loadDefaults() {
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
}

func InitConfig() error {
	viper.SetConfigName("server")             // name of config file (without extension)
	viper.AddConfigPath("/etc/gatekeeper/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.gatekeeper/") // call multiple times to add many search paths
	viper.AddConfigPath(".")                  // optionally look for config in the working directory
	loadDefaults()
	return viper.ReadInConfig()
}
