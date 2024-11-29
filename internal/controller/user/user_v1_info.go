package user

import (
	"context"
	"gf-admin/internal/logic/user"

	"gf-admin/api/user/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	return user.NewUserLogic().Info(ctx, req)
}
