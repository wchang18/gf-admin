package role

import (
	"context"
	"gf-admin/internal/logic/role"

	"gf-admin/api/role/v1"
)

func (c *ControllerV1) RoleList(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	return role.NewRoleLogic().RoleList(ctx, req)
}
