package kube_utils

import (
	"context"
	v1core "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/klog/v2"
	"kube-center/kube-utils/client"
	"os"
	"strings"
)

func (k *Pod) Exec(cmd []string) error {
	restconfig, err, coreclient := client.InitRestClient()
	if err != nil {
		return err
	}

	req := coreclient.RESTClient().Post().Namespace(k.Namespace).Resource("pods").Name(k.Name).
		SubResource("exec").VersionedParams(&v1core.PodExecOptions{
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       false,
		Container: k.ContainerName,
		Command:   cmd,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(restconfig, "POST", req.URL())
	if err != nil {
		klog.Error("ERROR：", err)
		return err
	}

	err = exec.StreamWithContext(context.TODO(), remotecommand.StreamOptions{
		Stdin:             strings.NewReader(""),
		Stdout:            os.Stdout,
		Stderr:            os.Stderr,
		Tty:               false,
		TerminalSizeQueue: nil,
	})

	if err != nil {
		return err
	}

	return err
}
