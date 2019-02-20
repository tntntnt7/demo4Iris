package controller

import (
	"github.com/kataras/iris"

	. "demo4Iris/service"
)

func GetUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	u := GetUserById(id)
	ctx.JSON(u)
}