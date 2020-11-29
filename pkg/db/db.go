package db

import (
	"database/sql"

	// import do driver Postgres
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// DB struct armazena a conexão com o banco
type DB struct {
	Client *sql.DB
}

// Get retorna a conexão com o banco
func Get(connStr string) (*DB, error) {
	db, err := get(connStr)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

// Close fecha a conexão com o banco
func (db *DB) Close() error {
	return db.Client.Close()
}

func get(connStr string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	return db, nil
}
