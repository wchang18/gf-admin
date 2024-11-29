package user

import (
	"context"
	"gf-admin/internal/logic/user"

	"gf-admin/api/user/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return user.NewUserLogic().Logout(ctx, req)
}
