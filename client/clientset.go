package client

import (
	"context"
	"flag"
	"fmt"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/utils/pointer"
	"path/filepath"
)

const (
	NAMESPACE      = "cto"
	DEPLOYMENTNAME = "kubecto-deploy"
	IMAGE          = "nginx:1.13"
	PORT           = 80
	REPLICAS       = 2
)

func main() {
	var kubeconfig *string
	// home是家目录，如果取得家目录的值，就可以用来做默认值
	if home := homedir.HomeDir(); home != "" {
		// 如果输入了kubeconfig参数，该参数的值就是kubeconfig文件的绝对路径
		// 如果没有输入kubeconfig的参数，就用默认的~/.kube/config
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file")
	} else {
		// 如果取不到当前用户的家目录，就没办法设置kubeconfig的默认目录，只能从入参中取
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	// 从本机加载kubeconfig配置文件，因此第一个参数为空字符串
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	// kubeconfig加载失败就直接退出了
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	createNamespace(clientset)
	createDeployment(clientset)
}

func createNamespace(clientset *kubernetes.Clientset) {
	namespaceClient := clientset.CoreV1().Namespaces()

	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: NAMESPACE,
		},
	}

	result, err := namespaceClient.Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Create namespace %s \n", result.GetName())
}

// 新建deployment
func createDeployment(clientset *kubernetes.Clientset) {
	// 如果希望在default命名空间下场景可以引用corev1.NamespaceDefault默认字符
	// deploymentClient := clientset.AppsV1().Deployments(corev1.NamespaceDefault)
	// 拿到deployment的客户端
	deploymentClient := clientset.AppsV1().Deployments(NAMESPACE)

	deployment := &appv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: DEPLOYMENTNAME,
		},
		Spec: appv1.DeploymentSpec{
			Replicas: pointer.Int32(REPLICAS),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "kubecto",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "kubecto",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "web",
							Image: IMAGE,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: PORT,
								},
							},
						},
					},
				},
			},
		},
	}

	// create Deployment
	fmt.Println("Creating deployment...")
	result, err := deploymentClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}