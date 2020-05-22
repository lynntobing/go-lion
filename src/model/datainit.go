/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:01
 * @File  : datainit.go
 * @Description: 数据初始化
 */
package model

import "fmt"

func dataInit() {
	var (
		tenantCount int
		menuCount   int
		roleCount   int
		err         error
	)
	if err = DB().Model(&Tenant{}).Count(&tenantCount).Error; err != nil {
		fmt.Println("tenant count:", err)
	}
	if err = DB().Model(&Menu{}).Count(&menuCount).Error; err != nil {
		fmt.Println("menu count:", err)
	}
	if err = DB().Model(&Role{}).Count(&roleCount).Error; err != nil {
		fmt.Println("role count:", err)
	}
	if tenantCount == 0 {
		tenantDataInit()
	}
	if menuCount == 0 {
		menuDataInit()
	}
	if roleCount == 0 {
		RoleDataInit()
	}
}

func tenantDataInit() {

}

func menuDataInit() {

}

func RoleDataInit() {

}
