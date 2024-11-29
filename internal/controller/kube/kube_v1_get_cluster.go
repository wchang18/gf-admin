package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) GetCluster(ctx context.Context, req *v1.GetClusterReq) (res *v1.GetClusterRes, err error) {
	return kube.GetClusterInfo(ctx, req)
}
