package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	client, error := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))

	if error != nil {
		log.Fatal("Error: ", error)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	error = client.Connect(ctx)

	if error != nil {
		log.Fatal(error)
	}

	err := client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connected")

	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	collection := client.Database("golangAPI").Collection(collectionName)
	return collection

}
