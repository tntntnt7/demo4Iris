package main

import (
	"github.com/kataras/iris"
	"gowork/common/config"
	. "gowork/routes"
)

func main() {
	App.Run(iris.Addr(":" + config.Config.Port), iris.WithConfiguration(iris.YAML("./common/config/iris.yml")))
}