package kube_utils

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube-center/kube-utils/client"
)

func (k *Deployment) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.AppsV1().Deployments(k.Namespace).Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *DaemonSet) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.AppsV1().DaemonSets(k.Namespace).Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *ReplicaSet) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.AppsV1().ReplicaSets(k.Namespace).Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *StatefulSet) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.AppsV1().StatefulSets(k.Namespace).Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *CronJob) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.BatchV1().CronJobs(k.Namespace).Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Job) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.BatchV1().Jobs(k.Namespace).Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *NameSpace) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Service) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.CoreV1().Services(k.Namespace).Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Ingress) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.NetworkingV1().Ingresses(k.Namespace).Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Pod) Delete(opt *metav1.DeleteOptions) error {
	var clientset, err = client.InitClient()
	if err != nil {
		return err
	}

	err = clientset.CoreV1().Pods(k.Namespace).Delete(context.TODO(), k.Name, *opt)
	if err != nil {
		return err
	}
	return err
}
