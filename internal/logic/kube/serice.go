package kube

import (
	"context"
	v1 "gf-admin/api/kube/v1"
	kv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateService(ctx context.Context, req *v1.CreateServiceReq) (res *v1.CreateServiceRes, err error) {
	apiObj, err := YamlStrToApiObject(req.Service)
	if err != nil {
		return
	}
	serviceObject := &kv1.Service{}
	serviceObject = apiObj.(*kv1.Service)
	_, err = DefaultCluster.ClientSet.CoreV1().Services(req.Namespace).Create(ctx, serviceObject, metav1.CreateOptions{})
	return
}

func DeleteService(ctx context.Context, req *v1.DeleteServiceReq) (res *v1.DeleteServiceRes, err error) {
	err = DefaultCluster.ClientSet.CoreV1().Services(req.Namespace).Delete(ctx, req.Name, metav1.DeleteOptions{})
	return
}

func GetServiceList(ctx context.Context, req *v1.GetServiceListReq) (res *v1.GetServiceListRes, err error) {
	res = &v1.GetServiceListRes{}
	serviceList, err := DefaultCluster.ClientSet.CoreV1().Services(req.Namespace).List(ctx, metav1.ListOptions{})

	list := make([]Service, 0)
	for _, service := range serviceList.Items {
		prots := make([]string, 0)
		for _, servicePort := range service.Spec.Ports {
			if service.Spec.Type == kv1.ServiceTypeNodePort {
				s := servicePort.TargetPort.StrVal + ":" + string(servicePort.NodePort) + "/" + string(servicePort.Protocol)
				prots = append(prots, s)
			} else {
				s := servicePort.TargetPort.StrVal + "/" + string(servicePort.Protocol)
				prots = append(prots, s)
			}
		}

		svc := Service{
			Name:        service.Name,
			Namespace:   service.Namespace,
			Type:        string(service.Spec.Type),
			ClusterIp:   service.Spec.ClusterIP,
			Ports:       prots,
			Selector:    service.Spec.Selector,
			Labels:      service.Labels,
			Annotations: service.Annotations,
		}
		list = append(list, svc)
	}
	res.List = list
	return
}
