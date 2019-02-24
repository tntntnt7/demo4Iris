package service

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	. "github.com/mongodb/mongo-go-driver/mongo"
	. "github.com/tntntnt7/demo4Iris/common/config"
	"github.com/tntntnt7/demo4Iris/common/utils"
	"log"
)

func UserSignUp(user *bson.M) interface{} {
	insertOneRes, err := userRep().InsertOne(utils.GetContext(), user)
	if err != nil { log.Fatal(err) }

	return insertOneRes.InsertedID
}

func Login(name, password string) (u bson.M) {
	err := userRep().FindOne(utils.GetContext(), bson.M{"Name": name, "Password": password}).Decode(&u)
	if err != nil { log.Fatal(err) }
	return
}

func userRep() *Collection {
	_rep := Mongo.Database("demo4Iris").Collection("user")
	return _rep
}

func GetUserById(id string) (u bson.M) {
	oid, _ := primitive.ObjectIDFromHex(id)
	
	err := userRep().FindOne(utils.GetContext(), bson.D{{"_id", oid}}).Decode(&u)
	if err != nil { log.Println(err) }
	
	return
}

func GetUsers() (result []bson.M) {
	result = []bson.M{}
	ctx := utils.GetContext()

	cur, err := userRep().Find(utils.GetContext(), bson.D{})
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var ret bson.M
		if err = cur.Decode(&ret); err != nil {
			log.Fatal(err)
		}
		result = append(result, ret)
	}
	if err := cur.Err(); err != nil { log.Fatal(err) }

	return
}

func UpdateUser(user bson.M) interface{} {
	var oid primitive.ObjectID
	id, ok := user["_id"].(string)
	if ok {
		oid, _ = primitive.ObjectIDFromHex(id)
	}

	updateOneRes, err := userRep().UpdateOne(
		utils.GetContext(),
		bson.M{"_id": oid},
		bson.M{"$set":
			bson.M{
				"Name": user["Name"],
				"Password": user["Password"],
				"Age": user["Age"],
				"Gender": user["Gender"],
			},
		},
	)
	if err != nil { log.Fatal(err) }

	return updateOneRes.ModifiedCount
}

func DeleteUserById(id string) interface{} {
	oid, _ := primitive.ObjectIDFromHex(id)

	deleteOneRes, err := userRep().DeleteOne(utils.GetContext(), bson.D{{"_id", oid}})
	if err != nil { log.Println(err) }

	return deleteOneRes.DeletedCount
}