// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRole is the golang structure for table user_role.
type UserRole struct {
	Id        uint        `json:"id"        orm:"id"        ` //
	Userid    int         `json:"userId"    orm:"userId"    ` // 用户Id
	Roleid    int         `json:"roleId"    orm:"roleId"    ` // 角色Id
	Createdat *gtime.Time `json:"createdAt" orm:"createdAt" ` // 创建时间
}
