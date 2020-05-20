package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config deepin tracker default configurations
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	LDAP struct {
		Host        string `yaml:"host"`
		Port        int    `yaml:"port"`
		Dn          string `yaml:"dn"`
		Password    string `yaml:"password"`
		UserSearch  string `yaml:"user_search"`
		GroupSearch string `yaml:"group_search"`
	} `yaml:"ldap"`
}

func GetConfig(filename string) (*Config)  {
	var  conf Config
	if content,err:=ioutil.ReadFile(filename);err != nil{
		panic(err)
	}else {
		if err=yaml.Unmarshal(content,&conf);err != nil{
			panic(err)
		}
	}
	return &conf
}