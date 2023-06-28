package kube_utils

import (
	"context"
	v1app "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
	v1networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *Deployment) Update(yaml []byte, opt *metav1.UpdateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	deployment := obj.(*v1app.Deployment)

	_, err = clientset.AppsV1().Deployments(k.Namespace).Update(context.TODO(), deployment, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *DaemonSet) Update(yaml []byte, opt *metav1.UpdateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	daemonset := obj.(*v1app.DaemonSet)

	_, err = clientset.AppsV1().DaemonSets(k.Namespace).Update(context.TODO(), daemonset, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *ReplicaSet) Update(yaml []byte, opt *metav1.UpdateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	rs := obj.(*v1app.ReplicaSet)

	_, err = clientset.AppsV1().ReplicaSets(k.Namespace).Update(context.TODO(), rs, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Service) Update(yaml []byte, opt *metav1.UpdateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	service := obj.(*v1core.Service)

	_, err = clientset.CoreV1().Services(k.Namespace).Update(context.TODO(), service, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Ingress) Update(yaml []byte, opt *metav1.UpdateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	ingress := obj.(*v1networking.Ingress)

	_, err = clientset.NetworkingV1().Ingresses(k.Namespace).Update(context.TODO(), ingress, *opt)
	if err != nil {
		return err
	}

	return err
}
