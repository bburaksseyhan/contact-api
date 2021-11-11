package server

import (
	"github.com/bburaksseyhan/contact-api/src/pkg/client/mongodb"
	"github.com/bburaksseyhan/contact-api/src/pkg/handler"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Init(url string) {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	client := mongodb.ConnectMongoDb(url)

	handler := handler.NewContactHandler(client)

	router.GET("/health", handler.HealthCheck)

	log.Info("port is :8080\n", url)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8080")
}
