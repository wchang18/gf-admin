package v1

import "github.com/gogf/gf/v2/frame/g"

type GetJobsReq struct {
	g.Meta    `path:"/jobs" method:"get" tags:"Job" summary:"Get job info"`
	Namespace string `json:"namespace"`
}

type GetJobsRes struct {
	List interface{} `json:"list"`
}

type CreateJobReq struct {
	g.Meta    `path:"/createJob" method:"post"`
	Namespace string `p:"namespace"`
	Job       string `p:"job"`
}

type CreateJobRes struct{}

type DeleteJobReq struct {
	g.Meta    `path:"/deleteJob" method:"delete"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type DeleteJobRes struct{}

type GetCronJobsReq struct {
	g.Meta    `path:"/cronjobs" method:"get" tags:"Job" summary:"Get cronjob info"`
	Namespace string `json:"namespace"`
}

type GetCronJobsRes struct {
	List interface{} `json:"list"`
}

type CreateCronJobReq struct {
	g.Meta    `path:"/createCronJob" method:"post"`
	Namespace string `p:"namespace"`
	CronJob   string `p:"cronJob"`
}

type CreateCronJobRes struct{}

type DeleteCronJobReq struct {
	g.Meta    `path:"/deleteCronJob" method:"delete"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type DeleteCronJobRes struct{}
