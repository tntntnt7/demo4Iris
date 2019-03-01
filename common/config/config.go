package config

import (
	"github.com/gomodule/redigo/redis"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/utils"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"gopkg.in/yaml.v2"

	. "github.com/tntntnt7/demo4Iris/common/utils"
)

var App *iris.Application

var Config = config{}

var Mongo *mongo.Client

var Cache redis.Conn

type config struct {
	Port string		`yaml:"Port"`

	Mongodb struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	}	`yaml:"Mongodb"`

	Redis struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	}	`yaml:"Redis"`
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
		Logger.Fatalf("config.yml err: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		Logger.Fatalf("Unmarshal: %v", err)
	}
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
	ctx := GetContext()
	Mongo, err = mongo.Connect(ctx, url)
	err = Mongo.Ping(ctx, readpref.Primary())
	if err != nil {
		Logger.Fatalf("!!! mongodb connect error: %v", err)
	} else {
		Logger.Info(url + "	connect connected successfully!")
	}
}

func InitRedis() {
	url := Config.Redis.Host + ":" + Config.Redis.Port

	var err error
	Cache, err = redis.Dial("tcp", url)
	if err != nil {
		Logger.Fatalf("!!! redis connect error: %v", err)
	} else {
		Logger.Info("redis://" + url + "	connect connected successfully!")
	}
}