package v1

import "github.com/gogf/gf/v2/frame/g"

type GetServiceListReq struct {
	g.Meta    `path:"/serviceList" method:"get"`
	Namespace string `json:"namespace"`
}

type GetServiceListRes struct {
	List interface{} `json:"list"`
}

type CreateServiceReq struct {
	g.Meta    `path:"/createService" method:"post"`
	Namespace string `p:"namespace"`
	Service   string `p:"service"`
}

type CreateServiceRes struct{}

type DeleteServiceReq struct {
	g.Meta    `path:"/deleteService" method:"delete"`
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

type DeleteServiceRes struct{}
