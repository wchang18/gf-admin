package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) CreateService(ctx context.Context, req *v1.CreateServiceReq) (res *v1.CreateServiceRes, err error) {
	return kube.CreateService(ctx, req)
}
