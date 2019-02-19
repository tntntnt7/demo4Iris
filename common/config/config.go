package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	Port	string	`yaml:"Port"`
}

var Config config

func init() {
	Config = config{}
	yamlFile, err := ioutil.ReadFile("./common/config/config.yml")
	if err != nil {
		fmt.Println("config.yml err:", err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println("conf>>>", Config)
}
