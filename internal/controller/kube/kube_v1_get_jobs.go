package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) GetJobs(ctx context.Context, req *v1.GetJobsReq) (res *v1.GetJobsRes, err error) {
	return kube.GetJobs(ctx, req)
}
