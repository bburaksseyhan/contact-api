package server

import (
	"net/http"

	"github.com/bburaksseyhan/contact-api/src/pkg/client/mongodb"
	"github.com/gin-gonic/gin"
)

func Init() {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	mongodb.ConnectMongoDb("mongodb://localhost:27017")

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to MongoDb Client Tutorial")
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8080")
}
