package config

import (
	"github.com/astaxie/beego/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var App *iris.Application

var Config = config{}

var PathConfig string

type config struct {
	App	app	`yaml:"App"`
}

type app struct {
	Port	string	`yaml:"Port"`
}

func InitConfig() {
	// 获取运行时绝对路径
	file := filepath.Dir(os.Args[0])
	runPath, _ := filepath.Abs(file)

	PathConfig = runPath + "/config.yml"
	if !utils.FileExists(PathConfig) {
		workPath, _ := os.Getwd() // 获取当前目录(即执行命令的目录, 和pwd命令类似)
		PathConfig = workPath + "/config.yml"
	}

	// 读取config.yml
	yamlFile, err := ioutil.ReadFile(PathConfig)
	if err != nil {
		log.Fatalf("config.yml err: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func InitApp() {
	App = iris.New()
	App.Logger().SetLevel("debug")
	App.Use(recover.New())
	App.Use(logger.New())
}
