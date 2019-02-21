package config

import (
	"github.com/astaxie/beego/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/mongodb/mongo-go-driver/mongo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var App *iris.Application

var Config = config{}

var Mongo *mongo.Client

type config struct {
	App	app	`yaml:"App"`
}

type app struct {
	Port	string	`yaml:"Port"`
}

func InitConfig() {
	// 获取运行时绝对路径
	runPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	configPath := runPath + "/config.yml"
	if !utils.FileExists(configPath) {
		workPath, _ := os.Getwd() // 获取当前目录(即执行命令的目录, 和pwd命令类似)
		configPath = workPath + "/config.yml"
	}

	// 读取config.yml
	yamlFile, err := ioutil.ReadFile(configPath)
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

func InitMongodb() {
	var err error
	Mongo, err = mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("!!! mongodb connect error: %v", err)
	} else {
		log.Println("mongodb://localhost:27017 connect connected successfully!")
	}
}