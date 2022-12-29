package database

import (
	"context"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindServerByServerID(id string) (class.Guilds, bool) {
	var result class.Guilds

	err := Guilds.FindOne(context.TODO(), bson.D{{Key: "server_id", Value: id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return class.Guilds{ServerID: "", CreatedAt: 0}, false
	}
	if err != nil {
		utils.HandleServerError(err)
		return class.Guilds{ServerID: "", CreatedAt: 0}, false
	}

	return class.Guilds{ServerID: result.ServerID, CreatedAt: result.CreatedAt}, true
}

func FindServerByID(id string) (class.Guilds, bool) {
	var result class.Guilds

	err := Guilds.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return class.Guilds{ServerID: "", CreatedAt: 0}, false
	}
	if err != nil {
		utils.HandleServerError(err)
		return class.Guilds{ServerID: "", CreatedAt: 0}, false
	}

	return class.Guilds{ServerID: result.ServerID, CreatedAt: result.CreatedAt}, true
}

func CreateGuild(id string) string {
	guild := class.NewGuild(id)

	res, err := Guilds.InsertOne(context.TODO(), guild)
	if err != nil {
		utils.HandleServerError(err)
	}

	return res.InsertedID.(primitive.ObjectID).String()
}

func UpdateGuild(id string, update *class.Guilds) bool {

	filter := bson.D{{Key: "server_id", Value: id}}

	_, err := Guilds.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		utils.HandleServerError(err)
	}

	return true
}
