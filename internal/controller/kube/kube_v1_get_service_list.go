package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) GetServiceList(ctx context.Context, req *v1.GetServiceListReq) (res *v1.GetServiceListRes, err error) {
	return kube.GetServiceList(ctx, req)
}
