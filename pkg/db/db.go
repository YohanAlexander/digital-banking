package db

import (
	"database/sql"

	// import do driver Postgres
	_ "github.com/lib/pq"
	"github.com/yohanalexander/desafio-banking-go/pkg/logger"
)

// DB armazena a conexão com o banco de dados
type DB struct {
	Client *sql.DB
}

// GetDB retorna a conexão com o banco de dados
func GetDB(connStr string) (*DB, error) {
	db, err := getDB(connStr)
	if err != nil {
		logger.Info.Fatal(err.Error())
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

// CloseDB fecha a conexão com o banco de dados
func (db *DB) CloseDB() error {
	return db.Client.Close()
}

// getDB estabelece a conexão com o banco de dados
func getDB(connStr string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Info.Fatal(err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Info.Fatal(err.Error())
		return nil, err
	}

	return db, nil
}
