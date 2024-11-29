package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) CreateDeployment(ctx context.Context, req *v1.CreateDeploymentReq) (res *v1.CreateDeploymentRes, err error) {
	return kube.CreateDeployment(ctx, req)
}
