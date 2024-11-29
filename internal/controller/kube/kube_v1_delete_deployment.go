package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) DeleteDeployment(ctx context.Context, req *v1.DeleteDeploymentReq) (res *v1.DeleteDeploymentRes, err error) {
	return kube.DeleteDeployment(ctx, req)
}
