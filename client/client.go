package client

import (
	"flag"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sync"
)

func InClusterRestClient() (*kubernetes.Clientset, error) {
	restconfig, err := rest.InClusterConfig()
	if err != nil {
		hlog.Error(err)
	}

	clientset, err := kubernetes.NewForConfig(restconfig)
	if err != nil {
		hlog.Error(err)
	}

	return clientset, err
}

func InitRestClient() (*rest.Config, error, *corev1client.CoreV1Client) {
	// Instantiate loader for kubeconfig file.
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	// Get a rest.Config from the kubeconfig file.  This will be passed into all
	// the client objects we create.
	restConfig, err := kubeconfig.ClientConfig()
	if err != nil {
		hlog.Error(err)
	}

	// Create a Kubernetes core/v1 client.
	coreClient, err := corev1client.NewForConfig(restConfig)
	if err != nil {
		hlog.Error(err)
	}
	return restConfig, err, coreClient
}

var once sync.Once
var clientset *kubernetes.Clientset

// 单例模式

func InitClient() (*kubernetes.Clientset, error) {
	once.Do(func() {
		hlog.Info("start doInit...")
		clientset, _ = doInit()
	})

	return clientset, nil
}

func doInit() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}
