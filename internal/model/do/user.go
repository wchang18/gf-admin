// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} //
	Username  interface{} // 用户名
	Password  interface{} // 密码
	Nickname  interface{} // 昵称
	Email     interface{} // 邮箱
	Phone     interface{} // 电话
	Status    interface{} // 1正常2禁用3注销
	Isadmin   interface{} // 管理员：1是
	Createdat *gtime.Time // 创建时间
	Updatedat *gtime.Time // 更新时间
}
