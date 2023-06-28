package kube_utils

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube-center/kube-utils/client"
)

func (k *Deployment) Scale(replicas int32) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	opts, err := clientset.AppsV1().Deployments(k.Namespace).GetScale(context.TODO(), k.Name, metav1.GetOptions{})
	opts.Spec.Replicas = replicas
	_, err = clientset.AppsV1().Deployments(k.Namespace).UpdateScale(context.TODO(), k.Name, opts, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return err
}

func (k *StatefulSet) Scale(replicas int32) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	opts, err := clientset.AppsV1().StatefulSets(k.Namespace).GetScale(context.TODO(), k.Name, metav1.GetOptions{})
	opts.Spec.Replicas = replicas
	_, err = clientset.AppsV1().StatefulSets(k.Namespace).UpdateScale(context.TODO(), k.Name, opts, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return err
}
