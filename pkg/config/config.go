package config

import (
	"fmt"

	"github.com/ibnumardini/wilayah-indonesia-api/pkg/env"
)

type serverConfig struct {
	AppName string
	AppHost string
	AppPort int
}

type dbConfig struct {
	Driver string
	Host   string
	Port   int
	Name   string
	User   string
	Pass   string
}

func (db *dbConfig) ConnStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.User, db.Pass, db.Host, db.Port, db.Name)
}

type config struct {
	Server serverConfig
	Db     dbConfig
}

var C config

func InitConfig() error {
	C = config{
		Server: serverConfig{
			AppName: env.Get("APP_NAME", "wilayah-indonesia-api"),
		},
		Db: dbConfig{
			Driver: env.Get("DB_DRIVER", "mysql"),
			Host:   env.Get("DB_HOST", "localhost"),
			Port:   env.GetAsInt("DB_PORT", 3306),
			Name:   env.Get("DB_NAME", "wilayah-indonesia-api"),
			User:   env.Get("DB_USER", "root"),
			Pass:   env.Get("DB_PASS", "password"),
		},
	}

	return nil
}
