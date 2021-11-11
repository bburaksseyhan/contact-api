package main

import (
	"os"

	"github.com/bburaksseyhan/contact-api/src/cmd/utils"
	"github.com/bburaksseyhan/contact-api/src/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	config := read()
	log.Info("Config.yml", config.Database.Url)

	mongoUri := os.Getenv("MONGODB_URL")
	serverPort := os.Getenv("SERVER_PORT")
	dbName := os.Getenv("DBNAME")
	collection := os.Getenv("COLLECTION")

	if mongoUri != "" {
		config.Database.Url = mongoUri
		config.Server.Port = serverPort
		config.Database.DbName = dbName
		config.Database.Collection = collection
	}

	log.Info("MONGODB_URL", mongoUri)

	server.Init(config)
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
