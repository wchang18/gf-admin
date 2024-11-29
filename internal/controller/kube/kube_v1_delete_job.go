package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) DeleteJob(ctx context.Context, req *v1.DeleteJobReq) (res *v1.DeleteJobRes, err error) {
	return kube.DeleteJob(ctx, req)
}
