package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) GetNodes(ctx context.Context, req *v1.GetNodesReq) (res *v1.GetNodesRes, err error) {
	return kube.GetNodeList(ctx, req)
}
