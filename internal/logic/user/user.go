package user

import (
	"context"
	"fmt"
	v1 "gf-admin/api/user/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/entity"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/thoas/go-funk"
)

type UserLogic struct {
}

func NewUserLogic() *UserLogic {
	return &UserLogic{}
}

func (l *UserLogic) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	var user entity.User
	user.Username = req.Username
	user.Password = gmd5.MustEncryptString(req.Password)
	user.Email = req.Email
	user.Nickname = req.Nickname
	user.Phone = req.Phone
	user.Status = 1
	id, err := dao.User.CreateUser(&user)

	for _, item := range req.Roles {
		role, _ := dao.Role.GetRoleByCode(item)
		roleUser := new(entity.UserRole)
		roleUser.Roleid = int(role.Id)
		roleUser.Userid = int(id)
		dao.UserRole.DB().Model(dao.UserRole.Table()).Data(roleUser).Save()
	}

	res = &v1.CreateUserRes{
		Id: uint(id),
	}
	return
}

func (l *UserLogic) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	user, err := dao.User.GetUserById(req.Id)
	if user == nil || user.Id == 0 {
		err = gerror.New("用户不存在")
		return
	}

	user.Email = req.Email
	user.Nickname = req.Nickname
	user.Phone = req.Phone
	user.Status = req.Status
	user.Updatedat = gtime.Now()
	err = dao.User.UpdateUser(user)

	dao.UserRole.DeleteByUserId(req.Id)
	for _, item := range req.Roles {
		role, _ := dao.Role.GetRoleByCode(item)
		roleUser := new(entity.UserRole)
		roleUser.Roleid = int(role.Id)
		roleUser.Userid = req.Id
		dao.UserRole.DB().Model(dao.UserRole.Table()).Data(roleUser).Save()
	}

	return
}

func (l *UserLogic) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	err = dao.User.DeleteUser(req.Id)
	dao.UserRole.DeleteByUserId(req.Id)
	return
}

func (l *UserLogic) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	res = &v1.UserListRes{}
	offset := (req.Page - 1) * req.PageSize
	var (
		list  []*v1.User
		total int
		ids   []int
	)

	model := dao.User.DB().Model(dao.User.Table())
	if req.Search != "" {
		model.WhereLike("username", "%"+req.Search+"%")
	}

	model.WhereNot("status", 3).
		Order(GetSort(req.Sort)).
		Offset(offset).
		Limit(req.PageSize).
		ScanAndCount(&list, &total, false)
	for _, item := range list {
		ids = append(ids, item.Id)
	}
	mCode, mName, _ := dao.UserRole.GetRolesByUserIds(ids)
	for i, item := range list {
		if roles, ok := mCode[item.Id]; ok {
			item.Roles = roles
		} else {
			item.Roles = []string{}
		}
		if rolenames, ok := mName[item.Id]; ok {
			item.RoleNames = rolenames
		} else {
			item.RoleNames = []string{}
		}
		// item.RoleNames = mName[item.Id]
		list[i] = item
	}

	res.List = list
	res.Total = total
	return
}

func GetSort(sort string) string {
	sortStr := "id desc"
	if sort != "" {
		sortStr = sort[1:]
		if sort[0] == '+' {
			sortStr += " asc"
		} else if sort[0] == '-' {
			sortStr += " desc"
		}
	}
	return sortStr
}

func (l *UserLogic) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	user, err := dao.User.GetUserByUsername(req.Username)
	if user == nil || user.Id == 0 {
		err = gerror.New("用户不存在")
		return
	}

	if user.Password != gmd5.MustEncryptString(req.Password) {
		err = gerror.New("密码错误")
		return
	}

	res = &v1.LoginRes{
		Token: l.GenToken(user),
	}
	return
}

func (l *UserLogic) GenToken(user *entity.User) (token string) {
	ctx := context.Background()
	token = guid.S()
	tokenKey := GetTokenKey(token)
	g.Redis().SetEX(ctx, tokenKey, user, 86400)
	return
}

func GetTokenKey(token string) string {
	return fmt.Sprintf("token:%s", token)
}

func (l *UserLogic) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	tokenKey := GetTokenKey(req.Token)
	g.Redis().Del(ctx, tokenKey)
	return
}

func (l *UserLogic) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	var user entity.User
	rdsRes, _ := g.Redis().Get(ctx, GetTokenKey(req.Token))
	err = rdsRes.Struct(&user)
	res = &v1.InfoRes{
		Name:    user.Username,
		Avatar:  "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif?imageView2/1/w/80/h/80",
		Roles:   []string{"admin"},
		IsAdmin: user.Isadmin,
	}
	res.Roles, res.Routes = l.GetUserRolesAndRoutes(int(user.Id))
	return
}

func (l *UserLogic) GetUserRolesAndRoutes(userId int) (roleName, routeName []string) {
	var roles []*entity.Role
	dao.Role.DB().Model(dao.Role.Table()+" as r ").
		LeftJoin(dao.UserRole.Table()+" as ur on r.id = ur.roleId").
		Fields("r.*").
		Where("ur.userId", userId).
		Where("r.status", 1).
		Scan(&roles)
	for _, item := range roles {
		roleName = append(roleName, item.Name)
		tmp, _ := gjson.DecodeToJson(item.Routename)
		routeName = append(routeName, tmp.Var().Strings()...)
	}
	roleName = funk.UniqString(roleName)
	routeName = funk.UniqString(routeName)
	return
}
