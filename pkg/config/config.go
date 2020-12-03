package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config armazena as variáveis de ambiente
type Config struct {
	tokenKey string
	apiPort  string
	dbUser   string
	dbPass   string
	dbHost   string
	dbPort   string
	dbName   string
	debug    string
}

// GetConfig captura os valores das variáveis de ambiente
func GetConfig() *Config {
	conf := &Config{}

	conf.debug = viper.GetString(`DEBUG_MODE`)
	conf.dbHost = viper.GetString(`POSTGRES_HOST`)
	conf.dbPort = viper.GetString(`POSTGRES_PORT`)
	conf.dbUser = viper.GetString(`POSTGRES_USER`)
	conf.dbPass = viper.GetString(`POSTGRES_PASSWORD`)
	conf.dbName = viper.GetString(`POSTGRES_DB`)
	conf.apiPort = viper.GetString(`SERVER_ADDRESS`)
	conf.tokenKey = viper.GetString(`TOKEN_KEY`)

	return conf
}

// GetDBConnStr retorna a string da conexão com DB formatada
func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

// getDBConnStr formata a string da conexão com DB
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

// GetAPIPort retorna a porta do servidor da API
func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}

// GetDebugMode retorna o valor do modo de debug
func (c *Config) GetDebugMode() string {
	return c.debug
}

// GetTokenKey retorna o valor da chave para gerar o token JWT
func (c *Config) GetTokenKey() string {
	return c.tokenKey
}
