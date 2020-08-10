package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

//Server is struct for https
type Server struct {
	Https        bool          `yaml:"Https"`
	HttpPort     int           `yaml:"HttpPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

//Database is struct for mysql
type Database struct {
	Type     string `yaml:"Type"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Name     string `yaml:"Name"`
}

//Redis is struct for redis
type Redis struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	Db       int    `yaml:"Db"`
}

//Etcd is struct for etcd
type Etcd struct {
	Endpoints []string `yaml:"Endpoints"`
}

//Elastic is elasticsearch struct
type Elastic struct {
	Addr     []string `yaml:"Addr"`
	Username string   `yaml:"Username"`
	Password string   `yaml:"Password"`
}

//System struct 全局
type System struct {
	RunMode  string   `yaml:"RunMode"`
	SiteName string   `yaml:"SiteName"`
	Server   Server   `yaml:"Server"`
	Database Database `yaml:"Database"`
	Redis    Redis    `yaml:"Redis"`
	Etcd     Etcd     `yaml:"Etcd"`
	Elastic  Elastic  `yaml:"Elastic"`
}

//Setting 声明为结构体实体
var Setting System

func init() {
	config, err := ioutil.ReadFile("./config/system.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &Setting)
}

//LoadConfig 返回结构体实体
func LoadConfig() System {
	return Setting
}
