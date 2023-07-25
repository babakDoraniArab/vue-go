package config

import (
	"log"

	"github.com/spf13/viper"
)

func Set() {
	//	func configSet() config.Config {
	// basic configuration for reading the config.yaml file from the config folder
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in
	viper.AddConfigPath("/app/config")
	// reading the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal("Error reading the config file -")
		}
	}

	// unmarshalling the config file into the config struct

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("Error unmarshalling the config file -")
	}

}
