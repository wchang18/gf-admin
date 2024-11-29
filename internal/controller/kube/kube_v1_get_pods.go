package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) GetPods(ctx context.Context, req *v1.GetPodsReq) (res *v1.GetPodsRes, err error) {
	return kube.GetPodList(ctx, req)
}
