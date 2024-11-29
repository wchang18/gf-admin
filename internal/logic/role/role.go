package role

import (
	"context"
	v1 "gf-admin/api/role/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/logic/user"
	"gf-admin/internal/model/entity"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type RoleLogic struct {
}

func NewRoleLogic() *RoleLogic {
	return &RoleLogic{}
}

func (l *RoleLogic) CreateRole(ctx context.Context, req *v1.CreateRoleReq) (res *v1.CreateRoleRes, err error) {
	if dao.Role.CheckRoleExist(req.Code, req.Name, false, 0) {
		err = gerror.New("角色已存在")
		return
	}

	var role entity.Role
	role.Code = req.Code
	role.Name = req.Name
	role.Remark = req.Remark
	// role.Routes = gjson.MustEncodeString(req.Routes)
	role.Routes, role.Routename = l.GetRoutes(req.Routes)
	role.Sort = req.Sort
	role.Status = req.Status
	id, err := dao.Role.CreateRole(&role)
	res = &v1.CreateRoleRes{
		Id: id,
	}
	return
}

func (l *RoleLogic) GetRoutes(route interface{}) (routeStr, routeName string) {
	routeStr = gjson.MustEncodeString(route)
	// routeMap := gconv.Map(route)
	names := make([]string, 0)
	for _, item := range gconv.SliceMap(route) {
		for key, value := range item {
			if key == "name" {
				names = append(names, value.(string))
			}
			if key == "children" {
				for _, it := range gconv.SliceMap(value) {
					for k, v := range it {
						if k == "name" {
							names = append(names, v.(string))
						}
					}
				}

			}
		}
	}

	routeName = gjson.MustEncodeString(names)
	return
}

func (l *RoleLogic) DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error) {
	err = dao.Role.DeleteRole(req.Id)
	return
}

func (l *RoleLogic) UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (res *v1.UpdateRoleRes, err error) {
	role, err := dao.Role.GetRoleById(req.Id)
	if role == nil || role.Id == 0 {
		err = gerror.New("角色不存在")
		return
	}

	if dao.Role.CheckRoleExist(req.Code, req.Name, true, req.Id) {
		err = gerror.New("角色名称或编码已存在")
		return
	}

	role.Code = req.Code
	role.Name = req.Name
	role.Remark = req.Remark
	// role.Routes = gjson.MustEncodeString(req.Routes)
	role.Routes, role.Routename = l.GetRoutes(req.Routes)
	role.Sort = req.Sort
	role.Status = req.Status
	role.Updatedat = gtime.Now()
	err = dao.Role.UpdateRole(role)
	return
}

func (l *RoleLogic) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	var data []*entity.Role
	res = &v1.RoleListRes{}
	total := 0
	offset := (req.Page - 1) * req.PageSize
	model := dao.Role.DB().Model(dao.Role.Table())
	if req.Search != "" {
		model = model.Where("name like ? or code like ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}
	model.Order(user.GetSort(req.Sort)).
		Offset(offset).
		Limit(req.PageSize).
		ScanAndCount(&data, &total, false)

	for _, v := range data {
		route, _ := gjson.Decode(v.Routes)
		role := v1.Role{
			Id:     int(v.Id),
			Code:   v.Code,
			Name:   v.Name,
			Remark: v.Remark,
			Routes: route,
			Sort:   v.Sort,
			Status: v.Status,
		}
		res.List = append(res.List, &role)
	}
	res.Total = total
	return
}

func (l *RoleLogic) GetAllRole(ctx context.Context, req *v1.GetAllRoleReq) (res *v1.
	GetAllRoleRes, err error) {
	res = &v1.GetAllRoleRes{}
	roleList, err := dao.Role.GetAllRole()
	for _, v := range roleList {
		role := v1.Role{
			Id:     int(v.Id),
			Code:   v.Code,
			Name:   v.Name,
			Remark: v.Remark,
			Sort:   v.Sort,
			Status: v.Status,
		}
		res.List = append(res.List, &role)
	}
	return
}
