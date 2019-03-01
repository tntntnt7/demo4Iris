package cache

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/mongodb/mongo-go-driver/bson"
	. "github.com/tntntnt7/demo4Iris/common/config"
	"github.com/tntntnt7/demo4Iris/common/utils"
)

func CGetUsers() interface{} {
	users, err := Cache.Do("LRANGE", "userList", 0, -1)
	keys, _ := redis.Strings(users, err)

	var ret []bson.M
	if len(keys) > 0 {
		for _, cell := range keys {
			retStrs, _ := redis.Strings(Cache.Do("HGETALL", cell))
			user := bson.M{}
			for i := 0; i < len(retStrs); i += 2 {
				user[retStrs[i]] = retStrs[i+1]
			}
			ret = append(ret, user)
		}

		return ret
	}
	return nil
}

func CSaveUsers(users []bson.M) {
	for _, user := range users {
		_, e1 := Cache.Do("HSET",
			user["_id"],
			"name", user["Name"],
			"password", user["Password"],
			"age", user["Age"],
			"gender", user["Gender"],
			)

		_, e2 := Cache.Do("LPUSH", "userList", user["_id"])

		_, e3 := Cache.Do("expire", user["Id"], 60)
		_, e4 := Cache.Do("expire", "userList", 60)

		if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
			utils.Logger.Error("存储失败!")
		}
	}
}

func SaveList(list []bson.M, key string) error {
	for _, cell := range list {
		_, e1 := Cache.Do("LPUSH", key, cell)
		_, e2 := Cache.Do("EXPIRE", key, 60)

		if e1 != nil || e2 != nil {
			utils.Logger.Error("存储失败!")
			return errors.New("存储失败!")
		}
	}
	return nil
}

func GetList(key string) interface{} {
	list, err := Cache.Do("LRANGE", key, 0, -1)
	if err != nil {
		return nil
	}
	return list
}