package user

import (
	"context"
	"gf-admin/internal/logic/user"

	"gf-admin/api/user/v1"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	return user.NewUserLogic().UserList(ctx, req)
}
