package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo.options"
	"log"
	"time"
)

func DBSetup() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongo://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

	defer cancel()

	err = client.Connect(context)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Failed to connect to mongoDb")
		return nil
	}

	fmt.Println("DB Connection Successful")
	return client
}

var Client *mongo.Client = DBSetup()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection{

}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {

}