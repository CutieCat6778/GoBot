package database

import (
	"context"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByMemberID(id string) (class.Members, bool) {
	var result class.Members

	err := Members.FindOne(context.TODO(), bson.D{{Key: "member_id", Value: id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		m := CreateMember(id)
		return class.Members{MemberID: m.MemberID, CreatedAt: m.CreatedAt, Tokens: m.Tokens, LastRefreshed: m.LastRefreshed}, true
	}
	if err != nil {
		utils.HandleServerError(err)
	}

	return class.Members{MemberID: result.MemberID, CreatedAt: result.CreatedAt, Tokens: result.Tokens, LastRefreshed: result.LastRefreshed}, true
}

func FindUserByID(id string) (class.Members, bool) {
	var result class.Members

	err := Members.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return class.Members{MemberID: "", CreatedAt: 0}, false
	}
	if err != nil {
		utils.HandleServerError(err)
	}

	return class.Members{MemberID: result.MemberID, CreatedAt: result.CreatedAt, Tokens: result.Tokens, LastRefreshed: result.LastRefreshed}, true
}

func CreateMember(id string) class.Members {
	guild := class.NewMember(id)

	_, err := Members.InsertOne(context.TODO(), guild)
	if err != nil {
		utils.HandleServerError(err)
	}

	return guild
}

func UpdateMember(id string, update *class.Members) bool {
	filter := bson.D{{Key: "member_id", Value: id}}

	_, err := Members.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		utils.HandleServerError(err)
	}

	return true
}

func RefreshToken(id string) bool {
	filter := bson.D{{Key: "member_id", Value: id}}
	currentTime := time.Now().Unix()

	m, f := FindUserByMemberID(id)
	if !f {
		utils.HandleServerError(errors.New("unable to find user " + id))
		return false
	}

	utils.HandleDebugMessage("Refreshing ", id, m.MemberID)

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "tokens", Value: 19}, {Key: "last_refreshed", Value: currentTime}}}}

	_, err := Members.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		utils.HandleServerError(err)
	}

	utils.HandleDebugMessage("Refreshed ", id)

	return true
}

func UserVoted(id string) bool {
	filter := bson.D{{Key: "member_id", Value: id}}
	currentTime := time.Now().Unix()

	m, f := FindUserByMemberID(id)
	if !f {
		utils.HandleServerError(errors.New("unable to find user " + id))
		return false
	}

	utils.HandleDebugMessage("Refreshing ", id, m.MemberID)

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "tokens", Value: m.Tokens + 20}, {Key: "last_refreshed", Value: currentTime}}}}

	_, err := Members.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		utils.HandleServerError(err)
	}

	utils.HandleDebugMessage("Refreshed ", id)

	return true
}

func RemoveToken(id string) (class.Members, bool) {
	filter := bson.D{{Key: "member_id", Value: id}}
	currentTime := time.Now().Unix()

	m, f := FindUserByMemberID(id)
	if !f {
		utils.HandleServerError(errors.New("unable to find user " + id))
	}

	tokenLeft := m.Tokens - 1

	interval := int64(1000 * 60 * 60 * 6)

	if m.Tokens == 0 && currentTime-m.LastRefreshed >= interval {
		refresh := RefreshToken(id)
		if !refresh {
			return m, false
		}
		return m, true
	} else if m.Tokens == 0 && currentTime-m.LastRefreshed < interval {
		return m, false
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "tokens", Value: tokenLeft}}}}

	_, err := Members.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		utils.HandleServerError(err)
	}

	return m, true
}
