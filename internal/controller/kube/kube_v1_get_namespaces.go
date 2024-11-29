package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) GetNamespaces(ctx context.Context, req *v1.GetNamespacesReq) (res *v1.GetNamespacesRes, err error) {
	return kube.GetNamespaceList(ctx, req)
}
