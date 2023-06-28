package kube_utils

import (
	"context"
	v1core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"kube-center/kube-utils/client"
	"sigs.k8s.io/yaml"
)

func (k *Pod) Describe() (v1core.Pod, v1core.EventList, error) {
	clientset, err := client.InitClient()
	if err != nil {
		return v1core.Pod{}, v1core.EventList{}, err
	}

	pod, err := clientset.CoreV1().Pods(k.Namespace).Get(context.TODO(), k.Name, metav1.GetOptions{})
	if err != nil {
		return v1core.Pod{}, v1core.EventList{}, err
	}

	eventList, err := clientset.CoreV1().Events(k.Namespace).List(context.TODO(), metav1.ListOptions{
		FieldSelector: "involvedObject.name=" + k.Name,
	})
	if err != nil {
		return v1core.Pod{}, v1core.EventList{}, err
	}

	return *pod, *eventList, err

}

func (k *Node) Describe() (v1core.Node, v1core.EventList, error) {
	clientset, err := client.InitClient()
	if err != nil {
		return v1core.Node{}, v1core.EventList{}, err
	}

	node, err := clientset.CoreV1().Nodes().Get(context.TODO(), k.Name, metav1.GetOptions{})
	if err != nil {
		return v1core.Node{}, v1core.EventList{}, err
	}

	// events
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	restconfig, err := kubeconfig.ClientConfig()
	if err != nil {
		return v1core.Node{}, v1core.EventList{}, err
	}

	coreclient, err := corev1client.NewForConfig(restconfig)
	if err != nil {
		return *node, v1core.EventList{}, err
	}

	req := coreclient.RESTClient().Get().Resource("events").Param("fieldSelector", "involvedObject.name="+k.Name)
	eventJson, _ := req.Do(context.TODO()).Raw()
	eventYaml, _ := yaml.JSONToYAML(eventJson)
	obj, _, err := scheme.Codecs.UniversalDeserializer().Decode(eventYaml, nil, nil)
	events := obj.(*v1core.EventList)
	if err != nil {
		return *node, v1core.EventList{}, err
	}

	return *node, *events, nil
}
