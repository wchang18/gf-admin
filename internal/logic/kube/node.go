package kube

import (
	"context"
	v1 "gf-admin/api/kube/v1"
	kv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeList(ctx context.Context, req *v1.GetNodesReq) (res *v1.GetNodesRes,
	err error) {
	res = &v1.GetNodesRes{}
	nodeList, err := DefaultCluster.ClientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	nodes := make([]Node, 0)
	for _, item := range nodeList.Items {
		node := Node{
			Name:                    item.Name,
			Labels:                  item.Labels,
			Annotations:             item.Annotations,
			CreationTimestamp:       item.CreationTimestamp.Time,
			Status:                  getReadyStatus(item.Status.Conditions),
			InternalIp:              getInternalIp(item.Status.Addresses),
			KernelVersion:           item.Status.NodeInfo.KernelVersion,
			KubeletVersion:          item.Status.NodeInfo.KubeletVersion,
			ContainerRuntimeVersion: item.Status.NodeInfo.ContainerRuntimeVersion,
			OsImage:                 item.Status.NodeInfo.OSImage}
		nodes = append(nodes, node)
	}
	res.List = nodes
	return
}

func getInternalIp(addresses []kv1.NodeAddress) string {
	for _, address := range addresses {
		if address.Type == kv1.NodeInternalIP {
			return address.Address
		}
	}
	return "notfound"
}

func getReadyStatus(conditions []kv1.NodeCondition) string {
	for _, condition := range conditions {
		if condition.Type == kv1.NodeReady {
			return string(condition.Status)
		}
	}
	return "notfound"
}
