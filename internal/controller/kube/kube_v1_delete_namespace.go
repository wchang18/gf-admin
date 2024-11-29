package kube

import (
	"context"
	"gf-admin/internal/logic/kube"

	"gf-admin/api/kube/v1"
)

func (c *ControllerV1) DeleteNamespace(ctx context.Context, req *v1.DeleteNamespaceReq) (res *v1.DeleteNamespaceRes, err error) {
	return kube.DeleteNamespace(ctx, req)
}
