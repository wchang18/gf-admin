package user

import (
	"context"
	"gf-admin/internal/logic/user"

	"gf-admin/api/user/v1"
)

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	return user.NewUserLogic().CreateUser(ctx, req)
}
