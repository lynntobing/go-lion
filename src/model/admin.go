/**
 * @Author: Lynn
 * @Time  : 2020-05-23 1:02
 * @File  : admin.go
 * @Description: ...
 */
package model

import (
	"github.com/jinzhu/gorm"
)

//租户
type Tenant struct {
	gorm.Model
	Tenancyname string //租户名[登录名]
	Name        string //租户名称
	Code        string //租户编码
	IsActive    bool   //是否可用
	IsTenant    bool   //是否是租主
}

//用户
type User struct {
	gorm.Model
	TenantId uint   //租户id
	UserName string //用户名
	Password string //密码
	Salt     string //密码盐值
	Name     string //姓名
	HeadImg  string //头像
	Phone    string //联系电话
	Email    string //邮箱
	IsActive bool   //是否可用
	IsAdmin  bool   //是否是管理员
}

//角色
type Role struct {
	gorm.Model
	TenantId    uint   //租户id
	Name        string //角色名称
	Code        string //角色编码
	Description string //角色描述
}

//用户角色
type UserRole struct {
	ID     uint //用户角色id
	UserId uint //用户id
	RoleId uint //角色id
}

//角色菜单表
type RoleMenu struct {
	ID     uint //角色菜单id
	RoleId uint //角色id
	MenuId uint //菜单id
}

//菜单
type Menu struct {
	gorm.Model
	Name      string //菜单名称
	Url       string //菜单url
	Icon      string //菜单icon
	Pid       uint   //父级菜单id
	IsDisplay bool   //是否展示
	Sort      int    //排序
}
