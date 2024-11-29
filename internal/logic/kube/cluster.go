package kube

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

type Cluster struct {
	Name          string
	Config        *rest.Config
	ClientSet     *kubernetes.Clientset
	DynamicClient dynamic.Interface
}

var DefaultCluster Cluster

const fixedConfigPath = "./manifest/config/.kube/default"

func SetDefaultCluster() {
	configFile := fixedConfigPath + "/config"
	_, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	KubeConfig, _ := clientcmd.BuildConfigFromFlags("", configFile)
	clientSet, err := kubernetes.NewForConfig(KubeConfig)
	if err != nil {
		panic(err)
	}
	dynamicClient, err := dynamic.NewForConfig(KubeConfig)
	if err != nil {
		panic(err)
	}
	DefaultCluster = Cluster{
		Name:          "default",
		Config:        KubeConfig,
		ClientSet:     clientSet,
		DynamicClient: dynamicClient,
	}
}
