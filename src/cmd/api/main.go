package main

import (
	"github.com/bburaksseyhan/contact-api/src/cmd/utils"
	"github.com/bburaksseyhan/contact-api/src/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	config := read()

	server.Init(config.Database.Url)
}

func read() utils.Configuration {
	//Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var config utils.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Error("Unable to decode into struct, %v", err)
	}

	return config
}
