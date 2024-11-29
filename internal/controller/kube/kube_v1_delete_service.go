package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) DeleteService(ctx context.Context, req *v1.DeleteServiceReq) (res *v1.DeleteServiceRes, err error) {
	return kube.DeleteService(ctx, req)
}
