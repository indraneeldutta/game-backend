package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// SetupEnvironment sets up the configs and environment for the application to start
func SetupEnvironment() error {
	err := viper.BindEnv("gopath")
	if err != nil {
		logrus.WithError(err).Fatal("gopath load failed")
	}

	viper.SetConfigName("config")

	viper.AddConfigPath("./resources/")
	viper.SetConfigType("yml")

	err = viper.ReadInConfig()
	if err != nil {
		logrus.WithError(err).Fatal("viper read failed")
	}

	return err

}
