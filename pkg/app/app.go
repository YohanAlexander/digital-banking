package app

import (
	"github.com/yohanalexander/desafio-banking-go/pkg/config"
	"github.com/yohanalexander/desafio-banking-go/pkg/db"
	"github.com/yohanalexander/desafio-banking-go/pkg/logger"
)

// App armazena configurações usadas em toda a API
type App struct {
	DB  *db.DB
	Cfg *config.Config
}

// GetApp captura variáveis de ambiente e conecta ao DB
func GetApp() (*App, error) {
	cfg := config.GetConfig()
	db, err := db.GetDB(cfg.GetDBConnStr())

	if err != nil {
		logger.Error.Fatal(err.Error())
		return nil, err
	}

	return &App{
		DB:  db,
		Cfg: cfg,
	}, nil
}
