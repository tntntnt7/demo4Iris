package user_service

import (
	"github.com/mongodb/mongo-go-driver/bson"

	. "github.com/tntntnt7/demo4Iris/cache/user_cache"
	. "github.com/tntntnt7/demo4Iris/common/utils"
	. "github.com/tntntnt7/demo4Iris/models/user"
)

func UserSignUp(user *bson.M) interface{} {
	return HandleSuccessfulResult(Add(user))
}

func Login(name, password string) interface{} {
	return HandleSuccessfulResult(Get(name, password))
}

func GetUserById(id string) interface{} {
	return HandleSuccessfulResult(GetOne(id))
}

func GetUsers() interface{} {
	// 读缓存
	cacheRet := GetList()
	if cacheRet != nil {
		return cacheRet
	}

	result := GetAll()

	// 写缓存
	SaveList(result)

	return HandleSuccessfulResult(result)
}

func UpdateUser(user bson.M) interface{} {
	return HandleSuccessfulResult(Update(user))
}

func DeleteUserById(id string) interface{} {
	return HandleSuccessfulResult(DeleteOne(id))
}