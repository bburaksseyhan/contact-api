package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactHandler interface {
	HealthCheck(*gin.Context)
}

type contactHandler struct {
	client *mongo.Client
}

func NewContactHandler(client *mongo.Client) ContactHandler {
	return &contactHandler{client: client}
}

func (ch *contactHandler) HealthCheck(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer ctxErr()

	if ctxErr != nil {
		logrus.Error("somethig wrong!!!", ctxErr)
	}

	if err := ch.client.Ping(ctx, nil); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "unhealty"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "pong"})
}
