package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"tratnik.net/client/internal/model"
)

func GetConfigFromFile(fileName string) *model.Config {
	if fileName == "" {
		fileName = "config"
	}
	viper.SetConfigName(fileName)
	viper.AddConfigPath("/etc/client/")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Fatal("Unable to read config")
	}

	config := &model.Config{}
	if err := viper.Unmarshal(&config); err != nil {
		logrus.WithError(err).Fatal("Unable to unmarshal config")
	}

	return config
}
