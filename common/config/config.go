package config

import (
	//"context"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	utils2 "github.com/tntntnt7/demo4Iris/common/utils"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	//"time"

	"github.com/astaxie/beego/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/mongodb/mongo-go-driver/mongo"
	"gopkg.in/yaml.v2"
)

var App *iris.Application

var Config = config{}

var Mongo *mongo.Client

type config struct {
	Port string		`yaml:"Port"`

	Mongodb struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	}	`yaml:"Mongodb"`
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

	log.Println("config => ", Config)
}

func InitApp() {
	App = iris.New()
	App.Logger().SetLevel("debug")
	App.Use(recover.New())
	App.Use(logger.New())
}

func InitMongodb() {
	url := "mongodb://" + Config.Mongodb.Host + ":" + Config.Mongodb.Port

	var err error
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := utils2.GetContext()
	Mongo, err = mongo.Connect(ctx, url)
	err = Mongo.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("!!! mongodb connect error: %v", err)
	} else {
		log.Println(url + "	connect connected successfully!")
	}
}
