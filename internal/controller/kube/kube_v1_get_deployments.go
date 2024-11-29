package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) GetDeployments(ctx context.Context, req *v1.GetDeploymentsReq) (res *v1.GetDeploymentsRes, err error) {
	return kube.GetDeploymentList(ctx, req)
}
