package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) GetCronJobs(ctx context.Context, req *v1.GetCronJobsReq) (res *v1.GetCronJobsRes, err error) {
	return kube.GetCronJobs(ctx, req)
}
