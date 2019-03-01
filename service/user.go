package service

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	. "github.com/mongodb/mongo-go-driver/mongo"
	"github.com/tntntnt7/demo4Iris/cache"
	. "github.com/tntntnt7/demo4Iris/common/config"
	. "github.com/tntntnt7/demo4Iris/common/utils"
)

var _userRep *Collection

func userRep() *Collection {
	if _userRep == nil {
		_userRep = Mongo.Database("demo4Iris").Collection("user")
	}

	return _userRep
}

func UserSignUp(user *bson.M) interface{} {
	insertOneRes, err := userRep().InsertOne(GetContext(), user)
	if err != nil {
		Logger.Error("UserSignUp Fail")
		return HandleErrorResult(err)
	}

	return HandleSuccessfulResult(insertOneRes.InsertedID)
}

func Login(name, password string) interface{} {
	var user bson.M
	err := userRep().FindOne(GetContext(), bson.M{"Name": name, "Password": password}).Decode(&user)
	if err != nil {
		Logger.Error("Login Fail")
		return HandleErrorResult(err)
	}

	return HandleSuccessfulResult(user)
}

func GetUserById(id string) interface{} {
	var user bson.M
	oid, _ := primitive.ObjectIDFromHex(id)
	
	err := userRep().FindOne(GetContext(), bson.D{{"_id", oid}}).Decode(&user)
	if err != nil {
		Logger.Error("GetUserById Fail, id = " + id)
		return HandleErrorResult(err)
	}
	
	return HandleSuccessfulResult(user)
}

func GetUsers() interface{} {
	var result []bson.M

	// 读缓存
	CRet := cache.CGetUsers()
	if CRet != nil {
		return CRet
	}

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
	if err := cur.Err(); err != nil {
		Logger.Error("GetUsers Fail")
		HandleErrorResult(err)
	}

	// 写缓存
	cache.CSaveUsers(result)

	return HandleSuccessfulResult(result)
}

func UpdateUser(user bson.M) interface{} {
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

	return HandleSuccessfulResult(updateOneRes.ModifiedCount)
}

func DeleteUserById(id string) interface{} {
	oid, _ := primitive.ObjectIDFromHex(id)

	deleteOneRes, err := userRep().DeleteOne(GetContext(), bson.D{{"_id", oid}})
	if err != nil {
		Logger.Error("DeleteUserById Fail, userId = " + id)
		HandleErrorResult(err)
	}

	return HandleSuccessfulResult(deleteOneRes.DeletedCount)
}