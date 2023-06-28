package kube_utils

import (
	"context"
	v1app "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"kube-center/kube-utils/client"
)

func (k *Deployment) Patch(data string, pt *types.PatchType, opt *metav1.PatchOptions) (v1app.Deployment, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return v1app.Deployment{}, err
	}

	deployment, err := clientset.AppsV1().Deployments(k.Namespace).Patch(context.TODO(), k.Name, *pt, []byte(data), *opt)
	if err != nil {
		return v1app.Deployment{}, err
	}

	return *deployment, err
}

func (k *DaemonSet) Patch(data string, pt *types.PatchType, opt *metav1.PatchOptions) (v1app.DaemonSet, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return v1app.DaemonSet{}, err
	}

	daemonset, err := clientset.AppsV1().DaemonSets(k.Namespace).Patch(context.TODO(), k.Name, *pt, []byte(data), *opt)
	if err != nil {
		return v1app.DaemonSet{}, err
	}

	return *daemonset, err
}

func (k *StatefulSet) Patch(data string, pt *types.PatchType, opt *metav1.PatchOptions) (v1app.StatefulSet, error) {
	var clientset, err = client.InitClient()
	if err != nil {
		return v1app.StatefulSet{}, err
	}

	statefulset, err := clientset.AppsV1().StatefulSets(k.Namespace).Patch(context.TODO(), k.Name, *pt, []byte(data), *opt)
	if err != nil {
		return v1app.StatefulSet{}, err
	}

	return *statefulset, err
}
