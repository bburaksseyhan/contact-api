package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Init(url string) {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	//mongodb.ConnectMongoDb(url)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to MongoDb Client Tutorial"})
	})

	log.Info("port is :8080", url)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8080")
}
