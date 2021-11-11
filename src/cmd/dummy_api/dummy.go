package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bburaksseyhan/contact-api/src/pkg/client/mongodb"
	"github.com/bburaksseyhan/contact-api/src/pkg/model"
	"github.com/sirupsen/logrus"
)

func main() {

	contactsJson, err := os.Open("contacts.json")

	if err != nil {
		logrus.Error("contact.json an error occurred", err)
	}

	defer contactsJson.Close()

	var contacts []model.Contact

	byteValue, _ := ioutil.ReadAll(contactsJson)

	//unmarshall data
	if err := json.Unmarshal(byteValue, &contacts); err != nil {
		logrus.Error("unmarshall an error occurred", err)
	}

	logrus.Info("Data\n", len(contacts))

	//import mongo client
	client := mongodb.ConnectMongoDb("mongodb://localhost:27017")
	logrus.Info(client)

	defer client.Disconnect(context.TODO())

	collection := client.Database("ContactsService").Collection("contacts")

	logrus.Warn("Total data count:", &contacts)

	for _, item := range contacts {
		collection.InsertOne(context.TODO(), item)
	}

	logrus.Info("Data import finished...")
}
