package main

import (
	"github.com/kataras/iris"
	"log"

	. "github.com/tntntnt7/demo4Iris/common/config"
	. "github.com/tntntnt7/demo4Iris/routes"
)

func init() {
	InitConfig()
	InitMongodb()
	InitApp()
	InitRoute()
}

func main() {
	log.Println("Port => ", Config.Port)
	App.Run(iris.Addr(":" + Config.Port))
}