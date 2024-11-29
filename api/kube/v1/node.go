package v1

import "github.com/gogf/gf/v2/frame/g"

type GetNodesReq struct {
	g.Meta `path:"/nodes" method:"get" tags:"Node" summary:"Get node info"`
}

type GetNodesRes struct {
	List interface{} `json:"list"`
}
