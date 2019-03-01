package user

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	. "github.com/mongodb/mongo-go-driver/mongo"
	. "github.com/tntntnt7/demo4Iris/common/config"
	. "github.com/tntntnt7/demo4Iris/common/utils"
)


type User struct {
	Id 			string	`bson:"_id"`
	Name		string	`bson:"Name"`
	Password	string	`bson:"Password"`
	Age 		int		`bson:"Age"`
	Gender		string	`bson:"Gender"`
}

var _userRep *Collection

func userRep() *Collection {
	if _userRep == nil {
		_userRep = Mongo.Database("demo4Iris").Collection("user_cache")
	}

	return _userRep
}

func Add(u *bson.M) interface{} {
	insertOneRes, err := userRep().InsertOne(GetContext(), u)
	if err != nil {
		Logger.Error("UserSignUp Fail")
		return HandleErrorResult(err)
	}

	return insertOneRes.InsertedID
}

func Get(name, password string) interface{} {
	var user bson.M
	err := userRep().FindOne(GetContext(), bson.M{"Name": name, "Password": password}).Decode(&user)
	if err != nil {
		Logger.Error("Login Fail")
		return HandleErrorResult(err)
	}

	return user
}

func GetOne(id string) interface{} {
	var user bson.M
	oid, _ := primitive.ObjectIDFromHex(id)

	err := userRep().FindOne(GetContext(), bson.D{{"_id", oid}}).Decode(&user)
	if err != nil {
		Logger.Error("GetUserById Fail, id = " + id)
		return HandleErrorResult(err)
	}

	return user
}

func GetAll() (result []bson.M) {
	ctx := GetContext()

	cur, err := userRep().Find(GetContext(), bson.D{})
	if err != nil {
		Logger.Error("GetUsers Fail")
		HandleErrorResult(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var ret bson.M
		if err = cur.Decode(&ret); err != nil {
			Logger.Error("GetUsers Fail")
			HandleErrorResult(err)
		}
		result = append(result, ret)
	}

	return result
}

func Update(user bson.M) interface{} {
	var oid primitive.ObjectID
	id, ok := user["_id"].(string)
	if ok {
		oid, _ = primitive.ObjectIDFromHex(id)
	}

	updateOneRes, err := userRep().UpdateOne(
		GetContext(),
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
	if err != nil {
		Logger.Error("UpdateUser Fail, userId = " + id)
		HandleErrorResult(err)
	}

	return updateOneRes.ModifiedCount
}

func DeleteOne(id string) interface{} {
	oid, _ := primitive.ObjectIDFromHex(id)

	deleteOneRes, err := userRep().DeleteOne(GetContext(), bson.D{{"_id", oid}})
	if err != nil {
		Logger.Error("DeleteUserById Fail, userId = " + id)
		HandleErrorResult(err)
	}

	return deleteOneRes.DeletedCount
}