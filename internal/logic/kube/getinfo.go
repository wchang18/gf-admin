package kube

import (
	"context"
	v1 "gf-admin/api/kube/v1"
	"math"

	// kv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetVersion(ctx context.Context, req *v1.GetVersionReq) (res *v1.GetVersionRes, err error) {
	res = &v1.GetVersionRes{}
	clientSet := DefaultCluster.ClientSet
	version, err := clientSet.ServerVersion()
	res.Version = version
	return
}

func GetClusterInfo(ctx context.Context, req *v1.GetClusterReq) (res *v1.GetClusterRes, err error) {
	extraClusterInfo := ClusterInfo{}

	clientSet := DefaultCluster.ClientSet

	nodeList, err := clientSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}
	var (
		TotalMemory float64
		UsedMemory  float64
	)
	nodes := nodeList.Items
	extraClusterInfo.TotalNodeNum = len(nodes)
	for i := range nodes {
		conditions := nodes[i].Status.Conditions
		for i := range conditions {
			if conditions[i].Type == "Ready" {
				if conditions[i].Status == "True" {
					extraClusterInfo.ReadyNodeNum += 1
				}
			}
		}
		cpu := nodes[i].Status.Allocatable.Cpu().AsApproximateFloat64()
		extraClusterInfo.TotalCpu += cpu
		memory := nodes[i].Status.Allocatable.Memory().AsApproximateFloat64()
		TotalMemory += memory
	}

	podsList, err := clientSet.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}
	pods := podsList.Items
	for i := range pods {
		for j := range pods[i].Spec.Containers {
			cpu := pods[i].Spec.Containers[j].Resources.Requests.Cpu().AsApproximateFloat64()
			extraClusterInfo.UsedCpu += cpu
			memory := pods[i].Spec.Containers[j].Resources.Requests.Memory().AsApproximateFloat64()
			UsedMemory += memory
		}
	}
	extraClusterInfo.UsedCpu = CovertRate(extraClusterInfo.UsedCpu)
	extraClusterInfo.UsedMemoryRate = CovertRate((UsedMemory / TotalMemory) * 100)
	extraClusterInfo.UsedMemory = UsedMemory
	extraClusterInfo.TotalMemory = TotalMemory
	res = &v1.GetClusterRes{
		Cluster: extraClusterInfo,
	}
	return
}

func CovertRate(number float64) float64 {
	return math.Round(number*100) / 100
}
