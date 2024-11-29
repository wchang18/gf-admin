package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) CreateJob(ctx context.Context, req *v1.CreateJobReq) (res *v1.CreateJobRes, err error) {
	return kube.CreateJob(ctx, req)
}
