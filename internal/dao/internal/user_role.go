// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserRoleDao is the data access object for table user_role.
type UserRoleDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserRoleColumns // columns contains all the column names of Table for convenient usage.
}

// UserRoleColumns defines and stores column names for table user_role.
type UserRoleColumns struct {
	Id        string //
	Userid    string // 用户Id
	Roleid    string // 角色Id
	Createdat string // 创建时间
}

// userRoleColumns holds the columns for table user_role.
var userRoleColumns = UserRoleColumns{
	Id:        "id",
	Userid:    "userId",
	Roleid:    "roleId",
	Createdat: "createdAt",
}

// NewUserRoleDao creates and returns a new DAO object for table data access.
func NewUserRoleDao() *UserRoleDao {
	return &UserRoleDao{
		group:   "default",
		table:   "user_role",
		columns: userRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserRoleDao) Columns() UserRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
