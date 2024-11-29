package role

import (
	"context"
	"gf-admin/internal/logic/role"

	"gf-admin/api/role/v1"
)

func (c *ControllerV1) DeleteRole(ctx context.Context, req *v1.DeleteRoleReq) (res *v1.DeleteRoleRes, err error) {
	return role.NewRoleLogic().DeleteRole(ctx, req)
}
