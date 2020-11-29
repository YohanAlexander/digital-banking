package app

import (
	"github.com/sirupsen/logrus"
	"github.com/yohanalexander/desafio-banking-go/pkg/config"
	"github.com/yohanalexander/desafio-banking-go/pkg/db"
)

// App encapsula a conexão com o banco e configurações
type App struct {
	DB  *db.DB
	Cfg *config.Config
}

// Get retorna struct App para a API
func Get() (*App, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	return &App{
		DB:  db,
		Cfg: cfg,
	}, nil
}
