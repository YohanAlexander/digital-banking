package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/yohanalexander/desafio-banking-go/pkg/config"
	"github.com/yohanalexander/desafio-banking-go/pkg/db"
)

// App armazena configurações usadas em toda a API
type App struct {
	DB  *db.DB
	Cfg *config.Config
	Vld *validator.Validate
	Log *logrus.Logger
}

// GetApp captura variáveis de ambiente e conecta ao DB
func GetApp() (*App, error) {
	log := logrus.New()
	vld := validator.New()
	cfg := config.GetConfig()
	db, err := db.GetDB(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}

	return &App{
		DB:  db,
		Cfg: cfg,
		Vld: vld,
		Log: log,
	}, nil
}
