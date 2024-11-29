package v1

import "github.com/gogf/gf/v2/frame/g"

type GetNamespacesReq struct {
	g.Meta `path:"/namespaces" method:"get" tags:"Namespace" summary:"Get namespace info"`
}

type GetNamespacesRes struct {
	List interface{} `json:"list"`
}

type CreateNamespaceReq struct {
	g.Meta `path:"/createNamespace" method:"post" tags:"Kubernetes" summary:"Create Kubernetes namespace"`
	Name   string `json:"name"`
}

type CreateNamespaceRes struct{}

type DeleteNamespaceReq struct {
	g.Meta `path:"/deleteNamespace" method:"delete" tags:"Kubernetes" `
	Name   string `json:"name"`
}

type DeleteNamespaceRes struct{}
