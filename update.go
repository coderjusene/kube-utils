package kube_utils

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/coderjusene/kube-utils/client"
	"strings"
)

func (k *Deployment) UpdateImage(tag string) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	opts, err := clientset.AppsV1().Deployments(k.Namespace).Get(context.TODO(), k.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	image := strings.Split(opts.Spec.Template.Spec.Containers[0].Image, ":")[0]
	newImage := strings.Join([]string{image, tag}, ":")
	opts.Spec.Template.Spec.Containers[0].Image = newImage
	_, err = clientset.AppsV1().Deployments(k.Namespace).Update(context.TODO(), opts, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return err
}

func (k *DaemonSet) UpdateImage(tag string) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	opts, err := clientset.AppsV1().DaemonSets(k.Namespace).Get(context.TODO(), k.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	image := strings.Split(opts.Spec.Template.Spec.Containers[0].Image, ":")[0]
	newImage := strings.Join([]string{image, tag}, ":")
	opts.Spec.Template.Spec.Containers[0].Image = newImage
	_, err = clientset.AppsV1().DaemonSets(k.Namespace).Update(context.TODO(), opts, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return err
}

func (k *StatefulSet) UpdateImage(tag string) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	opts, err := clientset.AppsV1().StatefulSets(k.Namespace).Get(context.TODO(), k.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	image := strings.Split(opts.Spec.Template.Spec.Containers[0].Image, ":")[0]
	newImage := strings.Join([]string{image, tag}, ":")
	opts.Spec.Template.Spec.Containers[0].Image = newImage
	_, err = clientset.AppsV1().StatefulSets(k.Namespace).Update(context.TODO(), opts, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return err
}
