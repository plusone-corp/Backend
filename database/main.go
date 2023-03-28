package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"plusone/backend/config"
	"time"
)

var (
	Client           *mongo.Client
	UserCollection   *mongo.Collection
	PostCollection   *mongo.Collection
	EventsCollection *mongo.Collection
	Context          context.Context
)

func init() {
	Client, err := mongo.NewClient(options.Client().ApplyURI(config.MONGO_URL))
	if err != nil {
		panic(err)
	}
	Context, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(Context)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database!")

	UserCollection = Client.Database("PlusOne").Collection("users")
	PostCollection = Client.Database("PlusOne").Collection("posts")
	EventsCollection = Client.Database("PlusOne").Collection("events")
}
