package kube

import (
	"context"
	v1 "gf-admin/api/kube/v1"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetJobs(ctx context.Context, req *v1.GetJobsReq) (res *v1.GetJobsRes, err error) {
	res = &v1.GetJobsRes{}
	data, err := DefaultCluster.ClientSet.BatchV1().Jobs(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}
	list := make([]Job, 0, len(data.Items))
	for _, item := range data.Items {
		job := Job{
			Name:              item.Name,
			Namespace:         item.Namespace,
			Annotations:       item.Annotations,
			CreationTimestamp: item.CreationTimestamp.Time,
			Labels:            item.Labels,
			Spec:              item.Spec,
			Status:            item.Status,
		}
		list = append(list, job)
	}
	res.List = list
	return
}

func CreateJob(ctx context.Context, req *v1.CreateJobReq) (res *v1.CreateJobRes, err error) {
	res = &v1.CreateJobRes{}
	job, err := YamlStrToApiObject(req.Job)
	if err != nil {
		return
	}
	_, err = DefaultCluster.ClientSet.BatchV1().Jobs(req.Namespace).Create(ctx, job.(*batchv1.Job), metav1.CreateOptions{})
	return
}

func DeleteJob(ctx context.Context, req *v1.DeleteJobReq) (res *v1.DeleteJobRes, err error) {
	res = &v1.DeleteJobRes{}
	err = DefaultCluster.ClientSet.BatchV1().Jobs(req.Namespace).Delete(ctx, req.Name, metav1.DeleteOptions{})
	return
}

func GetCronJobs(ctx context.Context, req *v1.GetCronJobsReq) (res *v1.GetCronJobsRes, err error) {
	res = &v1.GetCronJobsRes{}
	data, err := DefaultCluster.ClientSet.BatchV1().CronJobs(req.Namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}
	list := make([]CronJob, 0, len(data.Items))
	for _, item := range data.Items {
		cronJob := CronJob{
			Name:              item.Name,
			Namespace:         item.Namespace,
			Annotations:       item.Annotations,
			CreationTimestamp: item.CreationTimestamp.Time,
			Labels:            item.Labels,
			Spec:              item.Spec,
			Status:            item.Status,
		}
		list = append(list, cronJob)
	}
	res.List = list
	return
}

func CreateCronJob(ctx context.Context, req *v1.CreateCronJobReq) (res *v1.CreateCronJobRes, err error) {
	res = &v1.CreateCronJobRes{}
	job, err := YamlStrToApiObject(req.CronJob)
	if err != nil {
		return
	}
	_, err = DefaultCluster.ClientSet.BatchV1().CronJobs(req.Namespace).
		Create(ctx, job.(*batchv1.CronJob), metav1.CreateOptions{})
	return
}

func DeleteCronJob(ctx context.Context, req *v1.DeleteCronJobReq) (res *v1.DeleteCronJobRes, err error) {
	res = &v1.DeleteCronJobRes{}
	err = DefaultCluster.ClientSet.BatchV1().CronJobs(req.Namespace).Delete(ctx, req.Name, metav1.DeleteOptions{})
	return
}
