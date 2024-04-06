package databse

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
	mongoDb := os.Getenv("MONGODB_URL")

	Client, err := mongo.NewClient(options.Client().ApplyURI(mongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connected to mongoDb")
	return Client
}

var Client *mongo.Client = DbInstance()

func OpenCollection(Client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = Client.Database("cluster0").Collection(collectionName)
	return collection
}
