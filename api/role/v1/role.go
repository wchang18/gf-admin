package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateRoleReq struct {
	g.Meta `path:"/role/create" tags:"Role" method:"post" summary:"创建角色"`
	Role
}

type Role struct {
	Id     int         `json:"id"`
	Code   string      `json:"code" v:"required" `
	Name   string      `json:"name" v:"required" `
	Remark string      `json:"remark"`
	Sort   int         `json:"sort"`
	Status int         `json:"status" v:"required"`
	Routes interface{} `json:"routes"`
}

type CreateRoleRes struct {
	Id int64 `json:"id"`
}

type UpdateRoleReq struct {
	g.Meta `path:"/role/update" tags:"Role" method:"put" summary:"更新角色"`
	Role
}

type UpdateRoleRes struct {
}

type DeleteRoleReq struct {
	g.Meta `path:"/role/delete" tags:"Role" method:"delete" summary:"删除角色"`
	Id     int `json:"id" v:"required"`
}

type DeleteRoleRes struct {
}

type RoleListReq struct {
	g.Meta   `path:"/role/list" method:"get" summary:"获取角色列表"`
	Page     int    `p:"page" d:"1"`
	PageSize int    `p:"pageSize" d:"10"`
	Search   string `p:"search"`
	Sort     string `p:"sort"`
}

type RoleListRes struct {
	List  []*Role `json:"list"`
	Total int     `json:"total"`
}

type GetAllRoleReq struct {
	g.Meta `path:"/role/all" method:"get" summary:"获取所有角色"`
}

type GetAllRoleRes struct {
	List []*Role `json:"list"`
}
