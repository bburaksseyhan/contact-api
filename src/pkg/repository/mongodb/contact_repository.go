package repository

import (
	"context"

	"github.com/bburaksseyhan/contact-api/src/cmd/utils"
	"github.com/bburaksseyhan/contact-api/src/pkg/model"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ContactRepository interface {
	Get(ctx context.Context) ([]*model.Contact, error)
	GetContactByEmail(email string, ctx context.Context) (*model.Contact, error)
	Delete(id int, ctx context.Context) (int64, error)
}

type contactRepository struct {
	client *mongo.Client
	config *utils.Configuration
}

func NewContactRepository(config *utils.Configuration, client *mongo.Client) ContactRepository {
	return &contactRepository{config: config, client: client}
}

func (c *contactRepository) Get(ctx context.Context) ([]*model.Contact, error) {

	findOptions := options.Find()
	findOptions.SetLimit(100)

	var contacts []*model.Contact

	collection := c.client.Database(c.config.Database.DbName).Collection(c.config.Database.Collection)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem model.Contact
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
			return nil, err
		}

		contacts = append(contacts, &elem)
	}

	cur.Close(ctx)

	return contacts, nil
}

func (c *contactRepository) GetContactByEmail(email string, ctx context.Context) (*model.Contact, error) {

	findOptions := options.Find()
	findOptions.SetLimit(100)

	var contacts *model.Contact

	collection := c.client.Database(c.config.Database.DbName).Collection(c.config.Database.Collection)

	filter := bson.D{primitive.E{Key: "email", Value: email}}

	logrus.Info("Filter", filter)

	collection.FindOne(ctx, filter).Decode(&contacts)

	return contacts, nil
}

func (c *contactRepository) Delete(id int, ctx context.Context) (int64, error) {

	collection := c.client.Database(c.config.Database.DbName).Collection(c.config.Database.Collection)

	filter := bson.D{primitive.E{Key: "id", Value: id}}

	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)

		return 0, err
	}

	return deleteResult.DeletedCount, nil
}
