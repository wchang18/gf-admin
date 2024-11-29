package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"User" summary:"登录"`
	Username string `json:"username" v:"required"`
	Password string `json:"password" v:"required"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type InfoReq struct {
	g.Meta `path:"/info" method:"get" tags:"User" summary:"获取用户信息"`
	Token  string `json:"token"`
}

type InfoRes struct {
	Name    string   `json:"name"`
	Avatar  string   `json:"avatar"`
	Roles   []string `json:"roles"`
	Routes  []string `json:"routes"`
	IsAdmin int      `json:"isAdmin"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"User" summary:"登出"`
	Token  string `json:"token"`
}
type LogoutRes struct{}

type CreateUserReq struct {
	g.Meta `path:"/create" method:"post" tags:"User" summary:"创建用户"`
	User
}

type CreateUserRes struct {
	Id uint `json:"id"`
}

type UpdateUserReq struct {
	g.Meta `path:"/update" method:"put" tags:"User" summary:"更新用户"`
	User
}

type UpdateUserRes struct{}

type DeleteUserReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"User" summary:"删除用户"`
	Id     int `json:"id" v:"required"`
}

type DeleteUserRes struct{}

type User struct {
	Id        int      `json:"id"`
	Username  string   `json:"username" v:"required"`
	Password  string   `json:"password" v:"required"`
	Nickname  string   `json:"nickname" `
	Email     string   `json:"email" `
	Phone     string   `json:"phone" `
	Status    int      `json:"status"`
	Roles     []string `json:"roles"`
	RoleNames []string `json:"roleNames"`
}

type UserListReq struct {
	g.Meta   `path:"/list" method:"get" tags:"User" summary:"获取用户列表"`
	Search   string `p:"search"`
	Sort     string `p:"sort"`
	Page     int    `p:"page" d:"1"`
	PageSize int    `p:"pageSize" d:"10"`
}

type UserListRes struct {
	List  []*User `json:"list"`
	Total int     `json:"total"`
}
