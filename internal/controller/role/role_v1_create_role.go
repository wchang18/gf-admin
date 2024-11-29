package role

import (
	"context"
	"gf-admin/internal/logic/role"

	"gf-admin/api/role/v1"
)

func (c *ControllerV1) CreateRole(ctx context.Context, req *v1.CreateRoleReq) (res *v1.CreateRoleRes, err error) {
	return role.NewRoleLogic().CreateRole(ctx, req)
}
