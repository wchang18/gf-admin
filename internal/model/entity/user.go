// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint        `json:"id"        orm:"id"        ` //
	Username  string      `json:"username"  orm:"username"  ` // 用户名
	Password  string      `json:"password"  orm:"password"  ` // 密码
	Nickname  string      `json:"nickname"  orm:"nickname"  ` // 昵称
	Email     string      `json:"email"     orm:"email"     ` // 邮箱
	Phone     string      `json:"phone"     orm:"phone"     ` // 电话
	Status    int         `json:"status"    orm:"status"    ` // 1正常2禁用3注销
	Isadmin   int         `json:"isAdmin"   orm:"isAdmin"   ` // 管理员：1是
	Createdat *gtime.Time `json:"createdAt" orm:"createdAt" ` // 创建时间
	Updatedat *gtime.Time `json:"updatedAt" orm:"updatedAt" ` // 更新时间
}
