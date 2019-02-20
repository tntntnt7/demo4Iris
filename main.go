package main

import (
	"github.com/kataras/iris"

	. "demo4Iris/common/config"
	. "demo4Iris/routes"
)

func init() {
	InitApp()
	InitRoute()
}

func main() {
	App.Run(iris.Addr(":" + Config.Port))
}