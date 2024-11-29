package v1

import "github.com/gogf/gf/v2/frame/g"

type GetVersionReq struct {
	g.Meta `path:"/version" method:"get" tags:"Version" summary:"Get version info"`
}

type GetVersionRes struct {
	Version interface{} `json:"version"`
}

type GetClusterReq struct {
	g.Meta `path:"/cluster" method:"get" tags:"Version" summary:"Get cluster info"`
}

type GetClusterRes struct {
	Cluster interface{} `json:"cluster"`
}
