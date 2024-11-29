package role

import (
	"context"
	"gf-admin/internal/logic/role"

	"gf-admin/api/role/v1"
)

func (c *ControllerV1) GetAllRole(ctx context.Context, req *v1.GetAllRoleReq) (res *v1.GetAllRoleRes, err error) {
	return role.NewRoleLogic().GetAllRole(ctx, req)
}
