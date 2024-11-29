package kube

import (
	"context"
	v1 "gf-admin/api/kube/v1"
	kappv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodList(ctx context.Context, req *v1.GetPodsReq) (res *v1.GetPodsRes, err error) {
	podList, err := DefaultCluster.ClientSet.CoreV1().Pods(req.Namespace).List(ctx, metav1.ListOptions{})
	res = &v1.GetPodsRes{}
	list := make([]Pod, 0)
	if err != nil {
		return
	}
	for _, item := range podList.Items {
		containers := make([]Container, 0, len(item.Status.ContainerStatuses))
		for _, containerStatus := range item.Status.ContainerStatuses {
			container := Container{
				Name:         containerStatus.Name,
				Ready:        containerStatus.Ready,
				RestartCount: int(containerStatus.RestartCount),
				Image:        containerStatus.Image,
				ImageId:      containerStatus.ImageID, ContainerId: containerStatus.
						ContainerID,
			}
			containers = append(containers, container)
		}
		pod := Pod{
			Name:              item.Name,
			Namespace:         item.Namespace,
			Status:            string(item.Status.Phase),
			CreationTimestamp: item.CreationTimestamp.Time,
			Containers:        containers,
			Labels:            item.Labels,
			Annotations:       item.Annotations,
			PodIp:             item.Status.PodIP,
			NodeName:          item.Spec.NodeName,
		}
		list = append(list, pod)
	}
	res.List = list
	return
}

func GetDeploymentList(ctx context.Context, req *v1.GetDeploymentsReq) (res *v1.
	GetDeploymentsRes, err error) {
	res = &v1.GetDeploymentsRes{}
	list := make([]Deployment, 0)
	deployments, err := DefaultCluster.ClientSet.AppsV1().Deployments(req.Namespace).List(ctx,
		metav1.ListOptions{})
	if err != nil {
		return
	}
	for _, item := range deployments.Items {
		deployment := Deployment{
			Name:              item.Name,
			Namespace:         item.Namespace,
			Labels:            item.Labels,
			Annotations:       item.Annotations,
			CreationTimestamp: item.CreationTimestamp.Time,
			Replicas:          int(item.Status.Replicas),
			Status:            item.Status,
			Selector:          item.Spec.Selector.MatchLabels,
			Template:          item.Spec.Template,
		}
		list = append(list, deployment)
	}
	res.List = list
	return
}

func CreateDeployment(ctx context.Context, req *v1.CreateDeploymentReq) (res *v1.
	CreateDeploymentRes, err error) {
	res = &v1.CreateDeploymentRes{}
	deployment, err := YamlStrToApiObject(req.Deployment)
	if err != nil {
		return
	}
	_, err = DefaultCluster.ClientSet.AppsV1().Deployments(req.Namespace).Create(ctx,
		deployment.(*kappv1.Deployment), metav1.CreateOptions{})
	return
}

func DeleteDeployment(ctx context.Context, req *v1.DeleteDeploymentReq) (res *v1.DeleteDeploymentRes, err error) {
	res = &v1.DeleteDeploymentRes{}
	err = DefaultCluster.ClientSet.AppsV1().Deployments(req.Namespace).Delete(ctx, req.Name, metav1.DeleteOptions{})
	return
}
