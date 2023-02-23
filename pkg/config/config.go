package config

import (
	"log"

	"github.com/spf13/viper"
)

const configName = "config"
const configType = "yaml"

func Init(source string, c interface{}) {
	viper.AddConfigPath(source)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Panicln("Configuration file not found:", err)
		} else {
			log.Panicln("Failed to parse configuration file:", err)
		}
	}
	if err := viper.Unmarshal(c); err != nil {
		log.Panicln("Failed to unmarshal configuration file:", err)
	}
}
