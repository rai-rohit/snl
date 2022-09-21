package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DbConfig struct {
		Dbtype   string `yaml:"dbtype"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
	} `yaml:"database"`
}

func Load() (*Config, error) {
	var c Config
	yamlFile, err := ioutil.ReadFile("/home/rohit/go/src/snakes_ladders/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err :%v ", err)
		return &c,err 
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &c, err
}
