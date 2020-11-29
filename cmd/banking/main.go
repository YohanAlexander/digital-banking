package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
	"github.com/yohanalexander/desafio-banking-go/pkg/exit"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	//viper.SetConfigType("toml")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Using Config file: ", viper.ConfigFileUsed())
}

func main() {
	app, err := app.Get()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	exit.Init(func() {
		if err := app.DB.Close(); err != nil {
			logrus.Error(err.Error())
		}
	})
}
