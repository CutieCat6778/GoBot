package database

import (
	"context"
	"cutiecat6778/discordbot/class"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Guilds *mongo.Collection
)

func init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(class.DBKey))
	if err != nil {
		log.Fatal("Error while trying to connect to the database: ", err)
	}

	Guilds = client.Database("gobot").Collection("guild")

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		log.Fatal("Error while trying to disconnect to the database: ", err)
	// 		panic(err)
	// 	}
	// }()

}
