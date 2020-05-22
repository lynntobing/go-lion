/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:18
 * @File  : conf.go
 * @Description: ...
 */
package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type mysqlConf struct {
	Database string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}

type Conf struct {
	MysqlConf mysqlConf `yaml:"mysql"`
}

func InitConfig() (e *Conf, err error) {
	yamlFile, err := ioutil.ReadFile("./config/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Read err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &e)
	if err != nil {
		log.Fatal("Unmarshal:%v", err)
	}
	return e, err
}
