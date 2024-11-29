package kube

import (
	kappv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"time"
)

type ClusterInfo struct {
	UsedCpu        float64 `json:"used_cpu"`
	TotalCpu       float64 `json:"total_cpu"`
	UsedMemory     float64 `json:"used_memory"`
	TotalMemory    float64 `json:"total_memory"`
	UsedMemoryRate float64 `json:"used_memory_rate"`
	ReadyNodeNum   int     `json:"readyNodeNum"`
	TotalNodeNum   int     `json:"totalNodeNum"`
}

type Labels map[string]string
type Annotations map[string]string
type Selector map[string]string

type Node struct {
	Name                    string      `json:"name"`
	Status                  string      `json:"status"`
	Labels                  Labels      `json:"labels"`
	OsImage                 string      `json:"os_image"`
	InternalIp              string      `json:"internal_ip"`
	Annotations             Annotations `json:"annotations"`
	KernelVersion           string      `json:"kernel_version"`
	KubeletVersion          string      `json:"kubelet_version"`
	CreationTimestamp       time.Time   `json:"creation_timestamp"`
	ContainerRuntimeVersion string      `json:"container_runtime_version"`
}

type NameSpace struct {
	Name              string      `json:"name"`
	Status            string      `json:"status"`
	Labels            Labels      `json:"labels"`
	Annotations       Annotations `json:"annotations"`
	CreationTimestamp time.Time   `json:"creation_timestamp"`
}

type Service struct {
	Name        string      `json:"name"`
	Namespace   string      `json:"namespace"`
	Type        string      `json:"type"`
	ClusterIp   string      `json:"cluster_ip"`
	Ports       []string    `json:"ports"`
	Selector    Selector    `json:"selector"`
	Labels      Labels      `json:"labels"`
	Annotations Annotations `json:"annotations"`
	EndPoints   []string    `json:"end_points"`
}

type Pod struct {
	Name              string      `json:"name"`
	PodIp             string      `json:"pod_ip"`
	Status            string      `json:"status"`
	Labels            Labels      `json:"labels"`
	NodeName          string      `json:"node_name"`
	Namespace         string      `json:"namespace"`
	Containers        []Container `json:"containers"`
	Annotations       Annotations `json:"annotations"`
	CreationTimestamp time.Time   `json:"creation_timestamp"`
}

type Container struct {
	Name         string `json:"name"`
	Ready        bool   `json:"ready"`
	Image        string `json:"image"`
	ImageId      string `json:"image_id"`
	ContainerId  string `json:"container_id"`
	RestartCount int    `json:"restart_count"`
}

type Deployment struct {
	Name              string                  `json:"name"`
	Status            kappv1.DeploymentStatus `json:"status"`
	Labels            Labels                  `json:"labels"`
	Replicas          int                     `json:"replicas"`
	Namespace         string                  `json:"namespace"`
	Annotations       Annotations             `json:"annotations"`
	CreationTimestamp time.Time               `json:"creation_timestamp"`
	Selector          Labels                  `json:"selector"`
	Template          v1.PodTemplateSpec      `json:"template"`
}

type Job struct {
	Name              string            `json:"name"`
	Status            batchv1.JobStatus `json:"status"`
	Labels            Labels            `json:"labels"`
	Namespace         string            `json:"namespace"`
	Annotations       Annotations       `json:"annotations"`
	CreationTimestamp time.Time         `json:"creation_timestamp"`
	Spec              batchv1.JobSpec   `json:"spec"`
}

type CronJob struct {
	Name              string                `json:"name"`
	Status            batchv1.CronJobStatus `json:"status"`
	Labels            Labels                `json:"labels"`
	Namespace         string                `json:"namespace"`
	Annotations       Annotations           `json:"annotations"`
	CreationTimestamp time.Time             `json:"creation_timestamp"`
	Spec              batchv1.CronJobSpec   `json:"spec"`
}
