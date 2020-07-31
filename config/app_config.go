package config

import (
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	DbHost string `envconfig:"MYSQL_USERDB_HOST" required:"true"`
	DbPort int    `envconfig:"MYSQL_USERDB_PORT" required:"true"`
	DbName string `envconfig:"MYSQL_USERDB_DB" required:"true"`
	DbUser string `envconfig:"MYSQL_USERDB_USER" required:"true"`
	DbPass string `envconfig:"MYSQL_USERDB_PASS" required:"true"`
}

func GetAppConfig() AppConfig {
	var conf AppConfig
	envconfig.MustProcess("", &conf)
	return conf
}
