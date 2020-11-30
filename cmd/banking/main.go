package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/models/accounts"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/models/transfers"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/routers"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
	"github.com/yohanalexander/desafio-banking-go/pkg/exit"
	"github.com/yohanalexander/desafio-banking-go/pkg/logger"
	"github.com/yohanalexander/desafio-banking-go/pkg/server"
)

var api *app.App

func init() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warn("Falha ao carregar .env")
	} else {
		logrus.Info("Usando arquivo config:", viper.ConfigFileUsed())
		api, err = app.GetApp()
		if err != nil {
			logrus.Fatal(err.Error())
		} else {
			// migrando os schemas do DB
			err := api.DB.Client.AutoMigrate(&accounts.Account{}, &transfers.Transfer{})
			if err != nil {
				logrus.Fatal(err.Error())
			}
		}
	}
}

func main() {
	srv := server.
		GetServer().
		WithAddr(api.Cfg.GetAPIPort()).
		WithRouter(routers.GetRouter(api)).
		WithLogger(logger.Error)

	go func() {
		api.Log.Info("Iniciando servidor na porta", api.Cfg.GetAPIPort())
		if err := srv.StartServer(); err != nil {
			api.Log.Fatal(err.Error())
		}
	}()

	exit.Init(func() {
		if err := srv.CloseServer(); err != nil {
			api.Log.Error(err.Error())
		}

		if err := api.DB.CloseDB(); err != nil {
			api.Log.Error(err.Error())
		}
	})
}
