package user

import (
	"context"
	"gf-admin/internal/logic/user"

	"gf-admin/api/user/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return user.NewUserLogic().Login(ctx, req)
}
