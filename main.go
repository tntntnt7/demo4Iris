package main

import (
	"github.com/kataras/iris"
	. "github.com/tntntnt7/demo4Iris/common/config"
	. "github.com/tntntnt7/demo4Iris/common/utils"
	. "github.com/tntntnt7/demo4Iris/routes"
)

func init() {
	InitConfig()
	InitMongodb()
	InitApp()
	InitRoute()
	InitLogger()
}

func main() {
	Logger.Info("Server start at: " + Config.Port)
	App.Run(iris.Addr(":" + Config.Port))
}