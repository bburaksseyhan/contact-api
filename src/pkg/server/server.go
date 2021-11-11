package server

import (
	"github.com/bburaksseyhan/contact-api/src/cmd/utils"
	"github.com/bburaksseyhan/contact-api/src/pkg/client/mongodb"
	"github.com/bburaksseyhan/contact-api/src/pkg/handler"
	repository "github.com/bburaksseyhan/contact-api/src/pkg/repository/mongodb"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func Init(config utils.Configuration) {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	logrus.Info("Init\n %+d", config)
	client := mongodb.ConnectMongoDb(config.Database.Url)

	repo := repository.NewContactRepository(&config, client)
	handler := handler.NewContactHandler(client, repo)

	router.GET("/", handler.GetAllContacts)
	router.GET("/contacts/:email", handler.GetContactByEmail)
	router.POST("/contact/delete/:id", handler.DeleteContact)

	router.GET("/health", handler.HealthCheck)

	log.Info("port is :8080\n", config.Database.Url)

	// PORT environment variable was defined.
	router.Run(":" + config.Server.Port + "")
}
