package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) CreateCronJob(ctx context.Context, req *v1.CreateCronJobReq) (res *v1.CreateCronJobRes, err error) {
	return kube.CreateCronJob(ctx, req)
}
