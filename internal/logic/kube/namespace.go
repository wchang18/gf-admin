package kube

import (
	"context"
	v1 "gf-admin/api/kube/v1"
	kv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespaceList(ctx context.Context, req *v1.GetNamespacesReq) (res *v1.GetNamespacesRes, err error) {
	res = &v1.GetNamespacesRes{}
	clientSet := DefaultCluster.ClientSet
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}
	list := make([]NameSpace, 0)
	for _, namespace := range namespaceList.Items {
		var ns NameSpace
		ns.Name = namespace.Name
		ns.Labels = namespace.Labels
		ns.Annotations = namespace.Annotations
		ns.CreationTimestamp = namespace.CreationTimestamp.Time
		ns.Status = string(namespace.Status.Phase)
		list = append(list, ns)
	}
	res.List = list
	return
}

func CreateNamespace(ctx context.Context, req *v1.CreateNamespaceReq) (res *v1.CreateNamespaceRes, err error) {
	res = &v1.CreateNamespaceRes{}
	clientSet := DefaultCluster.ClientSet
	_, err = clientSet.CoreV1().Namespaces().Create(ctx, &kv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
	}, metav1.CreateOptions{})
	return
}

func DeleteNamespace(ctx context.Context, req *v1.DeleteNamespaceReq) (res *v1.DeleteNamespaceRes, err error) {
	res = &v1.DeleteNamespaceRes{}
	clientSet := DefaultCluster.ClientSet
	deletePolicy := metav1.DeletePropagationForeground
	err = clientSet.CoreV1().Namespaces().Delete(ctx, req.Name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	return
}
