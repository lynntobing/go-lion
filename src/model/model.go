/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:01
 * @File  : model.go
 * @Description: model
 */
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	conf "go-lion/conf"
)

var db *gorm.DB

func init() {
	config, _ := conf.InitConfig()
	var err error
	db, err = gorm.Open(config.MysqlConf.Name, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MysqlConf.User,
		config.MysqlConf.Password,
		config.MysqlConf.Host,
		config.MysqlConf.Port,
		config.MysqlConf.Database))
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	migration(db)

}

func DB() *gorm.DB {
	return db
}

func CloseDB() {
	defer db.Close()
}
