package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = connectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("Show").Collection(collectionName)
	return collection
}

func InsertCollection(client *mongo.Client, databaseName string, collectionName string, docs []interface{}) {

	collection := client.Database(databaseName).Collection(collectionName)
	insertManyResult, err := collection.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatal(err)
	}
	_ = insertManyResult
	//fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

}
