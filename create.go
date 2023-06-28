package kube_utils

import (
	"context"
	v1app "k8s.io/api/apps/v1"
	v2autoscaling "k8s.io/api/autoscaling/v2"
	v1batch "k8s.io/api/batch/v1"
	v1core "k8s.io/api/core/v1"
	v1networking "k8s.io/api/networking/v1"
	v1policy "k8s.io/api/policy/v1"
	v1rbac "k8s.io/api/rbac/v1"
	v1scheduling "k8s.io/api/scheduling/v1"
	v1storage "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"kube-center/kube-utils/client"
)

func initCreate(yaml []byte) (obj interface{}, clientset *kubernetes.Clientset, err error) {
	// yaml 序列化
	obj, _, err = scheme.Codecs.UniversalDeserializer().Decode(yaml, nil, nil)
	if err != nil {
		return
	}

	clientset, err = client.InitClient()
	if err != nil {
		return
	}

	return
}

func (k *Deployment) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	deployment := obj.(*v1app.Deployment)

	//// 校验typemeta
	//if deployment.Kind != "Deployment" && deployment.APIVersion != "apps/v1" {
	//	return errors.New(fmt.Sprintf("apiVersion: %s, kind: %s do not support",
	//		deployment.APIVersion, deployment.Kind))
	//}

	_, err = clientset.AppsV1().Deployments(k.Namespace).Create(context.TODO(), deployment, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *DaemonSet) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	daemonset := obj.(*v1app.DaemonSet)

	//// 校验typemeta
	//if daemonset.Kind != "DaemonSet" && daemonset.APIVersion != "apps/v1" {
	//	return errors.New(fmt.Sprintf("apiVersion: %s, kind: %s do not support",
	//		daemonset.APIVersion, daemonset.Kind))
	//}

	_, err = clientset.AppsV1().DaemonSets(k.Namespace).Create(context.TODO(), daemonset, *opt)
	if err != nil {
		return err
	}
	return err
}

func (k *ReplicaSet) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	replicaSet := obj.(*v1app.ReplicaSet)

	//// 校验typemeta
	//if replicaSet.Kind != "ReplicaSet" && replicaSet.APIVersion != "apps/v1" {
	//	return errors.New(fmt.Sprintf("apiVersion: %s, kind: %s do not support",
	//		replicaSet.APIVersion, replicaSet.Kind))
	//}

	_, err = clientset.AppsV1().ReplicaSets(k.Namespace).Create(context.TODO(), replicaSet, *opt)
	if err != nil {
		return err
	}
	return err
}

func (k *StatefulSet) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	statefulSet := obj.(*v1app.StatefulSet)

	_, err = clientset.AppsV1().StatefulSets(k.Namespace).Create(context.TODO(), statefulSet, *opt)
	if err != nil {
		return err
	}
	return err

}

func (k *CronJob) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	cronJob := obj.(*v1batch.CronJob)

	//// 校验typemeta
	//if cronJob.Kind != "CronJob" && cronJob.APIVersion != "batch/v1" {
	//	return errors.New(fmt.Sprintf("apiVersion: %s, kind: %s do not support",
	//		cronJob.APIVersion, cronJob.Kind))
	//}

	_, err = clientset.BatchV1().CronJobs(k.Namespace).Create(context.TODO(), cronJob, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Job) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	job := obj.(*v1batch.Job)

	_, err = clientset.BatchV1().Jobs(k.Namespace).Create(context.TODO(), job, *opt)
	if err != nil {
		return err
	}
	return err
}

func (k *Role) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	role := obj.(*v1rbac.Role)

	_, err = clientset.RbacV1().Roles(k.Namespace).Create(context.TODO(), role, *opt)
	if err != nil {
		return err
	}
	return err
}

func (k *RoleBinding) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	roleBinding := obj.(*v1rbac.RoleBinding)

	_, err = clientset.RbacV1().RoleBindings(k.Namespace).Create(context.TODO(), roleBinding, *opt)
	if err != nil {
		return err
	}
	return err
}

func (k *ClusterRole) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	clusterRole := obj.(*v1rbac.ClusterRole)
	_, err = clientset.RbacV1().ClusterRoles().Create(context.TODO(), clusterRole, *opt)
	if err != nil {
		return err
	}
	return err
}

func (k *ClusterRoleBinding) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	clusterRoleBinding := obj.(*v1rbac.ClusterRoleBinding)
	_, err = clientset.RbacV1().ClusterRoleBindings().Create(context.TODO(), clusterRoleBinding, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Ingress) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	ingress := obj.(*v1networking.Ingress)
	_, err = clientset.NetworkingV1().Ingresses(k.Namespace).Create(context.TODO(), ingress, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *HorizontalPodAutoscaler) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	hpa := obj.(*v2autoscaling.HorizontalPodAutoscaler)
	_, err = clientset.AutoscalingV2().HorizontalPodAutoscalers(k.Namespace).Create(context.TODO(), hpa, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *PodDisruptionBudget) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	pdb := obj.(*v1policy.PodDisruptionBudget)
	_, err = clientset.PolicyV1().PodDisruptionBudgets(k.Namespace).Create(context.TODO(), pdb, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Secret) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	secret := obj.(*v1core.Secret)
	_, err = clientset.CoreV1().Secrets(k.Namespace).Create(context.TODO(), secret, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Pod) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	pod := obj.(*v1core.Pod)
	_, err = clientset.CoreV1().Pods(k.Namespace).Create(context.TODO(), pod, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *PersistentVolume) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	pv := obj.(*v1core.PersistentVolume)
	_, err = clientset.CoreV1().PersistentVolumes().Create(context.TODO(), pv, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *Service) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	svc := obj.(*v1core.Service)
	_, err = clientset.CoreV1().Services(k.Namespace).Create(context.TODO(), svc, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *ServiceAccount) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	sa := obj.(*v1core.ServiceAccount)
	_, err = clientset.CoreV1().ServiceAccounts(k.Namespace).Create(context.TODO(), sa, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *NameSpace) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	ns := obj.(*v1core.Namespace)
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), ns, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *ConfigMap) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	cm := obj.(*v1core.ConfigMap)
	_, err = clientset.CoreV1().ConfigMaps(k.Namespace).Create(context.TODO(), cm, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *PersistentVolumeClaim) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	pvc := obj.(*v1core.PersistentVolumeClaim)
	_, err = clientset.CoreV1().PersistentVolumeClaims(k.Namespace).Create(context.TODO(), pvc, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *PriorityClass) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	priorityClass := obj.(*v1scheduling.PriorityClass)
	_, err = clientset.SchedulingV1().PriorityClasses().Create(context.TODO(), priorityClass, *opt)
	if err != nil {
		return err
	}

	return err
}

func (k *StorageClass) Create(yaml []byte, opt *metav1.CreateOptions) error {
	obj, clientset, err := initCreate(yaml)
	if err != nil {
		return err
	}

	storageClass := obj.(*v1storage.StorageClass)
	_, err = clientset.StorageV1().StorageClasses().Create(context.TODO(), storageClass, *opt)
	if err != nil {
		return err
	}

	return err
}
