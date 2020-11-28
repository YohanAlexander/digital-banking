package banking

import (
	"database/sql"
	"fmt"

	// import do driver Mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitDB abre a conexao com o banco de dados
func InitDB() (db *sql.DB) {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open(`mysql`, url)
	if err != nil {
		logrus.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}
