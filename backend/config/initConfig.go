package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workdir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read config file")
	}
}
