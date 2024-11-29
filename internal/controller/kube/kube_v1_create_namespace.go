package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) CreateNamespace(ctx context.Context, req *v1.CreateNamespaceReq) (res *v1.CreateNamespaceRes, err error) {
	return kube.CreateNamespace(ctx, req)
}
