package db

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB armazena a conexão com o banco de dados
type DB struct {
	Client *gorm.DB
}

// GetDB retorna a conexão com o banco de dados
func GetDB(connStr, debugMode string) (*DB, error) {
	db, err := getDB(connStr, debugMode)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

// CloseDB fecha a conexão com o banco de dados
func (db *DB) CloseDB() error {
	sqlDB, err := db.Client.DB()
	if err != nil {
		return errors.New("Erro ao fechar conexão com o banco")
	}
	return sqlDB.Close()
}

// getDB estabelece a conexão com o banco de dados
func getDB(connStr, debugMode string) (*gorm.DB, error) {

	var (
		db  *gorm.DB
		err error
	)

	if debugMode == "true" {
		if db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}); err != nil {
			return nil, errors.New("Erro ao abrir conexão com o banco")
		}

	} else {
		if db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{}); err != nil {
			return nil, errors.New("Erro ao abrir conexão com o banco")
		}
	}

	if sqlDB, err := db.DB(); err == nil {
		if err := sqlDB.Ping(); err != nil {
			return nil, errors.New("Erro ao abrir conexão com o banco")
		}
	} else {
		return nil, errors.New("Erro ao abrir conexão com o banco")
	}

	return db, nil
}
