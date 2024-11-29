// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package kube

import (
	"context"

	"gf-admin/api/kube/v1"
)

type IKubeV1 interface {
	GetJobs(ctx context.Context, req *v1.GetJobsReq) (res *v1.GetJobsRes, err error)
	CreateJob(ctx context.Context, req *v1.CreateJobReq) (res *v1.CreateJobRes, err error)
	DeleteJob(ctx context.Context, req *v1.DeleteJobReq) (res *v1.DeleteJobRes, err error)
	GetCronJobs(ctx context.Context, req *v1.GetCronJobsReq) (res *v1.GetCronJobsRes, err error)
	CreateCronJob(ctx context.Context, req *v1.CreateCronJobReq) (res *v1.CreateCronJobRes, err error)
	DeleteCronJob(ctx context.Context, req *v1.DeleteCronJobReq) (res *v1.DeleteCronJobRes, err error)
	GetNamespaces(ctx context.Context, req *v1.GetNamespacesReq) (res *v1.GetNamespacesRes, err error)
	CreateNamespace(ctx context.Context, req *v1.CreateNamespaceReq) (res *v1.CreateNamespaceRes, err error)
	DeleteNamespace(ctx context.Context, req *v1.DeleteNamespaceReq) (res *v1.DeleteNamespaceRes, err error)
	GetNodes(ctx context.Context, req *v1.GetNodesReq) (res *v1.GetNodesRes, err error)
	GetPods(ctx context.Context, req *v1.GetPodsReq) (res *v1.GetPodsRes, err error)
	GetDeployments(ctx context.Context, req *v1.GetDeploymentsReq) (res *v1.GetDeploymentsRes, err error)
	CreateDeployment(ctx context.Context, req *v1.CreateDeploymentReq) (res *v1.CreateDeploymentRes, err error)
	DeleteDeployment(ctx context.Context, req *v1.DeleteDeploymentReq) (res *v1.DeleteDeploymentRes, err error)
	GetServiceList(ctx context.Context, req *v1.GetServiceListReq) (res *v1.GetServiceListRes, err error)
	CreateService(ctx context.Context, req *v1.CreateServiceReq) (res *v1.CreateServiceRes, err error)
	DeleteService(ctx context.Context, req *v1.DeleteServiceReq) (res *v1.DeleteServiceRes, err error)
	GetVersion(ctx context.Context, req *v1.GetVersionReq) (res *v1.GetVersionRes, err error)
	GetCluster(ctx context.Context, req *v1.GetClusterReq) (res *v1.GetClusterRes, err error)
}
