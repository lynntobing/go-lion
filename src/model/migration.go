/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:01
 * @File  : migration.go
 * @Description: 数据库迁移
 */
package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func migration(db *gorm.DB) {
	tableList := []interface{}{
		&Tenant{}, &User{}, &Role{}, &UserRole{}, &RoleMenu{}, &Menu{},
	}
	num := 0
	for _, item := range tableList {
		if !db.HasTable(item) {
			db.AutoMigrate(item)
			num++
		}
	}
	fmt.Println("[data migration] success, migration table num:", num)
}
