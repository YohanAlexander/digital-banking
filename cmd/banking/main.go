package main

import (
	"github.com/spf13/viper"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/routers"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
	"github.com/yohanalexander/desafio-banking-go/pkg/exit"
	"github.com/yohanalexander/desafio-banking-go/pkg/logger"
	"github.com/yohanalexander/desafio-banking-go/pkg/server"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Info.Fatal(err.Error())
	}
	logger.Info.Println("Using Config file: ", viper.ConfigFileUsed())
}

func main() {
	app, err := app.GetApp()
	if err != nil {
		logger.Info.Fatal(err.Error())
	}

	srv := server.
		GetServer().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(routers.GetRouter(app)).
		WithLogger(logger.Error)

	go func() {
		logger.Info.Println("Starting server at ", app.Cfg.GetAPIPort())
		if err := srv.StartServer(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()

	exit.Init(func() {
		if err := srv.CloseServer(); err != nil {
			logger.Error.Println(err.Error())
		}

		if err := app.DB.CloseDB(); err != nil {
			logger.Error.Println(err.Error())
		}
	})
}
