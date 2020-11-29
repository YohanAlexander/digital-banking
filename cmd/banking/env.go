package banking

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// InitENV inicializa as variaveis de ambiente
func InitENV() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	viper.SetConfigType("toml")
	viper.SetConfigFile(".config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Using Config file: ", viper.ConfigFileUsed())

	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Warn("Banking service is Running in Debug Mode")
		return
	}
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Warn("Banking service is Running in Production Mode")
}
