package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"gowork/controller"
)

var App *iris.Application

func init()  {
	App = iris.New()
	App.Logger().SetLevel("debug")
	App.Use(recover.New())
	App.Use(logger.New())

	initRoute()
}

func initRoute() {
	user := App.Party("user")
	user.Get("/", controller.GetUser)
}
