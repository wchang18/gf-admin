package role

import (
	"context"
	"gf-admin/internal/logic/role"

	"gf-admin/api/role/v1"
)

func (c *ControllerV1) UpdateRole(ctx context.Context, req *v1.UpdateRoleReq) (res *v1.UpdateRoleRes, err error) {
	return role.NewRoleLogic().UpdateRole(ctx, req)
}
