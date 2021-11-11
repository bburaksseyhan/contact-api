package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bburaksseyhan/contact-api/src/pkg/model"
	db "github.com/bburaksseyhan/contact-api/src/pkg/repository/mongodb"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactHandler interface {
	GetAllContacts(*gin.Context)
	GetContactByEmail(*gin.Context)
	DeleteContact(*gin.Context)

	HealthCheck(*gin.Context)
}

type contactHandler struct {
	client     *mongo.Client
	repository db.ContactRepository
}

func NewContactHandler(client *mongo.Client, repo db.ContactRepository) ContactHandler {
	return &contactHandler{client: client, repository: repo}
}

func (ch *contactHandler) GetAllContacts(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer ctxErr()

	var contactList []*model.Contact

	//request on repository
	if result, err := ch.repository.Get(ctx); err != nil {
		logrus.Error(err)
	} else {
		contactList = result
	}

	c.JSON(http.StatusOK, gin.H{"contacts": &contactList})
}

func (ch *contactHandler) GetContactByEmail(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer ctxErr()

	var contactList *model.Contact

	//get parameter
	email := c.Param("email")

	//request on repository
	if result, err := ch.repository.GetContactByEmail(email, ctx); err != nil {
		logrus.Error(err)
	} else {
		contactList = result
	}

	c.JSON(http.StatusOK, gin.H{"contacts": contactList})
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

func (ch *contactHandler) DeleteContact(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer ctxErr()

	//get parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error("Can not convert to id", err)
	}

	//request on repository
	result, err := ch.repository.Delete(id, ctx)
	if err != nil {
		logrus.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"deleteResult.DeletedCount": result})
}
