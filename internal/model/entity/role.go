// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	Id        uint        `json:"id"        orm:"id"        ` //
	Code      string      `json:"code"      orm:"code"      ` // 角色编码
	Name      string      `json:"name"      orm:"name"      ` // 角色名
	Routename string      `json:"routeName" orm:"routeName" ` // 路由名称
	Routes    string      `json:"routes"    orm:"routes"    ` // 路由权限
	Remark    string      `json:"remark"    orm:"remark"    ` // 备注
	Sort      int         `json:"sort"      orm:"sort"      ` // 排序
	Status    int         `json:"status"    orm:"status"    ` // 1正常2禁用
	Createdat *gtime.Time `json:"createdAt" orm:"createdAt" ` // 创建时间
	Updatedat *gtime.Time `json:"updatedAt" orm:"updatedAt" ` // 更新时间
}
