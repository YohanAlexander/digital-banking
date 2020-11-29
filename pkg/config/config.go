package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config struct encapsula as variaveis de ambiente
type Config struct {
	address string
	dbUser  string
	dbPass  string
	dbHost  string
	dbPort  string
	dbName  string
	debug   string
}

// Get obtem os valores das variaveis de ambiente
func Get() *Config {
	conf := &Config{}

	conf.debug = viper.GetString(`DEBUG_MODE`)
	conf.dbHost = viper.GetString(`POSTGRES_HOST`)
	conf.dbPort = viper.GetString(`POSTGRES_PORT`)
	conf.dbUser = viper.GetString(`POSTGRES_USER`)
	conf.dbPass = viper.GetString(`POSTGRES_PASSWORD`)
	conf.dbName = viper.GetString(`POSTGRES_DB`)
	conf.address = viper.GetString(`SERVER_ADDRESS`)

	if conf.debug == "true" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Warn("Banking service is Running in Debug Mode")
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warn("Banking service is Running in Production Mode")
	}
	return conf
}

// GetDBConnStr formata a string da conexão
func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPass,
		dbhost,
		c.dbPort,
		dbname,
	)
}
