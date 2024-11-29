package v1

import "github.com/gogf/gf/v2/frame/g"

type GetPodsReq struct {
	g.Meta    `path:"/pods" method:"get" tags:"Pod" summary:"Get pod info"`
	Namespace string `json:"namespace"`
}

type GetPodsRes struct {
	List interface{} `json:"list"`
}

type GetDeploymentsReq struct {
	g.Meta    `path:"/deployments" method:"get" tags:"Deployment" summary:"Get deployment info"`
	Namespace string `json:"namespace"`
}

type GetDeploymentsRes struct {
	List interface{} `json:"list"`
}

type CreateDeploymentReq struct {
	g.Meta     `path:"/createDeployment" method:"post"`
	Namespace  string `p:"namespace"`
	Deployment string `p:"deployment"`
}

type CreateDeploymentRes struct {
}

type DeleteDeploymentReq struct {
	g.Meta    `path:"/deleteDeployment" method:"delete"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type DeleteDeploymentRes struct {
}
