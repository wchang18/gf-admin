// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-admin/internal/dao/internal"
)

// internalUserRoleDao is internal type for wrapping internal DAO implements.
type internalUserRoleDao = *internal.UserRoleDao

// userRoleDao is the data access object for table user_role.
// You can define custom methods on it to extend its functionality as you wish.
type userRoleDao struct {
	internalUserRoleDao
}

var (
	// UserRole is globally public accessible object for table user_role operations.
	UserRole = userRoleDao{
		internal.NewUserRoleDao(),
	}
)

// Fill with you ideas below.

func (dao *userRoleDao) GetRolesByUserIds(userIds []int) (map[int][]string,
	map[int][]string, error) {
	var (
		roles []*UserAndRole
		mCode = make(map[int][]string)
		mName = make(map[int][]string)
	)
	err := dao.DB().Model(dao.Table()).
		LeftJoin("role", "role.id=user_role.roleId").
		Fields("role.code,role.name,user_role.userId").
		Where("user_role.userId", userIds).
		Scan(&roles)
	for _, item := range roles {
		mCode[item.UserId] = append(mCode[item.UserId], item.RoleCode)
		mName[item.UserId] = append(mName[item.UserId], item.RoleName)
	}
	return mCode, mName, err
}

type UserAndRole struct {
	RoleCode string `orm:"code"`
	RoleName string `orm:"name"`
	UserId   int    `orm:"userId"`
}

func (dao *userRoleDao) DeleteByUserId(userId int) (err error) {
	_, err = dao.DB().Model(dao.Table()).Where("userId", userId).Delete()
	return
}
