package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Https        bool          `yaml:"Https"`
	HttpPort     int           `yaml:"HttpPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}
type Database struct {
	Type     string `yaml:"Type"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Name     string `yaml:"Name"`
}
type Redis struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"Password"`
	Db       int    `yaml:"Db"`
}
type System struct {
	RunMode  string   `yaml:"RunMode"`
	SiteName string   `yaml:"SiteName"`
	Server   Server   `yaml:"Server"`
	Database Database `yaml:"Database"`
	Redis    Redis    `yaml:"Redis"`
}

var Setting System

func init() {
	config, err := ioutil.ReadFile("./config/system.yaml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &Setting)
}
func LoadConfig() System {
	return Setting
}
