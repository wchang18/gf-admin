package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) GetVersion(ctx context.Context, req *v1.GetVersionReq) (res *v1.GetVersionRes, err error) {
	return kube.GetVersion(ctx, req)
}
