package app

import (
	"fmt"

	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	br_translations "github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/sirupsen/logrus"
	"github.com/yohanalexander/desafio-banking-go/pkg/config"
	"github.com/yohanalexander/desafio-banking-go/pkg/db"
)

// App armazena configurações usadas em toda a API
type App struct {
	DB    *db.DB
	Cfg   *config.Config
	Vld   *validator.Validate
	Log   *logrus.Logger
	Trans ut.Translator
}

// TranslateErrors traduz os erros de formatos JSON inválidos
func (app *App) TranslateErrors(err error) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(app.Trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

// GetApp captura variáveis de ambiente e conecta ao DB
func GetApp() (*App, error) {
	// definindo default logger
	log := logrus.New()
	// definindo validator de erros no formato JSON
	vld := validator.New()
	br := pt_BR.New()
	uni := ut.New(br, br)
	trans, _ := uni.GetTranslator("pt_BR")
	_ = br_translations.RegisterDefaultTranslations(vld, trans)
	// definindo configurações de ambiente
	cfg := config.GetConfig()
	// definindo conexão com o banco de dados
	db, err := db.GetDB(cfg.GetDBConnStr(), cfg.GetDebugMode())
	if err != nil {
		return nil, err
	}

	return &App{
		DB:    db,
		Cfg:   cfg,
		Vld:   vld,
		Log:   log,
		Trans: trans,
	}, nil
}
