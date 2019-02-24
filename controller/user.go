package controller

import (
	"github.com/kataras/iris"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/tntntnt7/demo4Iris/service"
)

func Register(ctx iris.Context) {
	user := &bson.M{}
	ctx.ReadJSON(user)

	ctx.JSON(service.UserSignUp(user))
}

func Login(ctx iris.Context) {
	data := &struct {
		Name string
		Password string
	}{}
	ctx.ReadJSON(data)

	ctx.JSON(service.Login(data.Name, data.Password))
}

func GetUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	user := service.GetUserById(id)

	ctx.JSON(user)
}

func GetAllUsers(ctx iris.Context) {
	ret := service.GetUsers()

	ctx.JSON(ret)
}

func UpdateUser(ctx iris.Context) {
	user := bson.M{}
	ctx.ReadJSON(&user)

	ctx.JSON(service.UpdateUser(user))
}

func DeleteUser(ctx iris.Context) {
	id := ctx.Params().Get("id")

	ctx.JSON(service.DeleteUserById(id))
}
