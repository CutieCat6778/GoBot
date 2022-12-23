package database

import (
	"context"
	"cutiecat6778/discordbot/class"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Guilds  *mongo.Collection
	Members *mongo.Collection
)

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(class.DBKey))
	if err != nil {
		log.Fatal("Error while trying to connect to the database: ", err)
	}

	Guilds = client.Database("gobot").Collection("guild")
	Members = client.Database("gobot").Collection("member")

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
}
