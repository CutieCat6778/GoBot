package database

import (
	"context"
	"cutiecat6778/discordbot/class"
	"cutiecat6778/discordbot/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByMemberID(id string) (class.Members, bool) {
	var result class.Members

	err := Members.FindOne(context.TODO(), bson.D{{Key: "member_id", Value: id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		var u class.Members
		var f bool
		m := CreateMember(id)
		if len(m) > 0 {
			u, f = FindUserByID(m)
			if !f {
				utils.SendErrorMessage("Problem while trying to fetch member data: ", err.Error())
				log.Fatal("Problem while trying to fetch data: ", err)
			}
		} else {
			utils.SendErrorMessage("Problem while trying to write member data: ", err.Error())
			log.Fatal("Problem while trying to write data: ", err)
		}
		return class.Members{MemberID: u.MemberID, CreatedAt: u.CreatedAt, Tokens: u.Tokens, LastRefreshed: u.LastRefreshed}, true
	}
	if err != nil {
		utils.SendErrorMessage("Problem while trying to fetch member data: ", err.Error())
		log.Fatal("Problem while trying to fetch data: ", err)
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
		utils.SendErrorMessage("Problem while trying to fetch member data: ", err.Error())
		log.Fatal("Problem while trying to fetch data: ", err)
	}

	return class.Members{MemberID: result.MemberID, CreatedAt: result.CreatedAt, Tokens: result.Tokens, LastRefreshed: result.LastRefreshed}, true
}

func CreateMember(id string) string {
	guild := class.NewMember(id)

	res, err := Members.InsertOne(context.TODO(), guild)
	if err != nil {
		log.Fatal("Problem while trying to write datas: ", err)
	}

	return res.InsertedID.(primitive.ObjectID).String()
}

func UpdateMember(id string, update *class.Members) bool {
	filter := bson.D{{Key: "member_id", Value: id}}

	_, err := Members.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		utils.SendErrorMessage("Failed to update Member! ", err.Error())
		log.Fatal("Failed to update Member: ", err)
	}

	return true
}

func RefreshToken(id string) bool {
	filter := bson.D{{Key: "member_id", Value: id}}
	current_time := time.Now().Unix()

	m, f := FindUserByMemberID(id)
	if !f {
		utils.SendErrorMessage("Failed to find user", "")
		log.Fatal("Failed to find user")
		return false
	}

	log.Println("Refreshing ", id, m.MemberID)

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "tokens", Value: 19}, {Key: "last_refreshed", Value: current_time}}}}

	_, err := Members.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		utils.SendErrorMessage("Failed to update member! ", err.Error())
		log.Fatal("Failed to update Member: ", err)
	}

	log.Println("Refreshed ", id)

	return true
}

func RemoveToken(id string) (class.Members, bool) {
	filter := bson.D{{Key: "member_id", Value: id}}
	current_time := time.Now().Unix()

	m, f := FindUserByMemberID(id)
	if !f {
		utils.SendErrorMessage("Failed to find user", "")
		log.Fatal("Failed to find user")
	}

	tokenLeft := m.Tokens - 1

	interval := int64(1000 * 60 * 60 * 6)

	if m.Tokens == 0 && current_time-m.LastRefreshed >= interval {
		refresh := RefreshToken(id)
		if !refresh {
			return m, false
		}
		return m, true
	} else if m.Tokens == 0 && current_time-m.LastRefreshed < interval {
		return m, false
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "tokens", Value: tokenLeft}}}}

	_, err := Members.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		utils.SendErrorMessage("Failed to update guild! ", err.Error())
		log.Fatal("Failed to update Member: ", err)
	}

	return m, true
}
