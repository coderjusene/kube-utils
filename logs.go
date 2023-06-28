package kube_utils

import (
	"bytes"
	"context"
	"io"
	v1core "k8s.io/api/core/v1"
	"github.com/coderjusene/kube-utils/client"
)

func (k *Pod) Logs(podLogOpts *v1core.PodLogOptions) (string, error) {
	podLogOpts.Container = k.ContainerName
	clientset, err := client.InitClient()
	if err != nil {
		return "", err
	}
	req := clientset.CoreV1().Pods(k.Namespace).GetLogs(k.Name, podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", err
	}

	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", err
	}

	str := buf.String()
	return str, nil
}
