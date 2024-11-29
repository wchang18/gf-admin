// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table role for DAO operations like Where/Data.
type Role struct {
	g.Meta    `orm:"table:role, do:true"`
	Id        interface{} //
	Code      interface{} // 角色编码
	Name      interface{} // 角色名
	Routename interface{} // 路由名称
	Routes    interface{} // 路由权限
	Remark    interface{} // 备注
	Sort      interface{} // 排序
	Status    interface{} // 1正常2禁用
	Createdat *gtime.Time // 创建时间
	Updatedat *gtime.Time // 更新时间
}
