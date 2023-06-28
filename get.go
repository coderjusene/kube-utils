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
	"github.com/coderjusene/kube-utils/client"
)

func (k *Deployment) Get(opt *metav1.GetOptions) (v1app.Deployment, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1app.Deployment{}, err
	}

	deployment, err := clientSet.AppsV1().Deployments(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1app.Deployment{}, err
	}

	return *deployment, nil
}

func (k *DaemonSet) Get(opt *metav1.GetOptions) (v1app.DaemonSet, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1app.DaemonSet{}, err
	}

	daemonset, err := clientSet.AppsV1().DaemonSets(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1app.DaemonSet{}, err
	}

	return *daemonset, err
}

func (k *ReplicaSet) Get(opt *metav1.GetOptions) (v1app.ReplicaSet, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1app.ReplicaSet{}, err
	}

	replicaset, err := clientSet.AppsV1().ReplicaSets(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1app.ReplicaSet{}, err
	}

	return *replicaset, err
}

func (k *StatefulSet) Get(opt *metav1.GetOptions) (v1app.StatefulSet, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1app.StatefulSet{}, err
	}

	statefulset, err := clientSet.AppsV1().StatefulSets(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1app.StatefulSet{}, err
	}

	return *statefulset, err
}

func (k *CronJob) Get(opt *metav1.GetOptions) (v1batch.CronJob, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1batch.CronJob{}, err
	}

	cronjob, err := clientSet.BatchV1().CronJobs(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1batch.CronJob{}, err
	}

	return *cronjob, err
}

func (k *Job) Get(opt *metav1.GetOptions) (v1batch.Job, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1batch.Job{}, err
	}

	job, err := clientSet.BatchV1().Jobs(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1batch.Job{}, err
	}

	return *job, err
}

func (k *Role) Get(opt *metav1.GetOptions) (v1rbac.Role, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1rbac.Role{}, err
	}

	role, err := clientSet.RbacV1().Roles(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1rbac.Role{}, err
	}

	return *role, err
}

func (k *RoleBinding) Get(opt *metav1.GetOptions) (v1rbac.RoleBinding, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1rbac.RoleBinding{}, err
	}

	rolebinding, err := clientSet.RbacV1().RoleBindings(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1rbac.RoleBinding{}, err
	}

	return *rolebinding, err
}

func (k *ClusterRole) Get(opt *metav1.GetOptions) (v1rbac.ClusterRole, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1rbac.ClusterRole{}, err
	}

	clusterrole, err := clientSet.RbacV1().ClusterRoles().Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1rbac.ClusterRole{}, err
	}

	return *clusterrole, err
}

func (k *ClusterRoleBinding) Get(opt *metav1.GetOptions) (v1rbac.ClusterRoleBinding, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1rbac.ClusterRoleBinding{}, err
	}

	clusterrolebinding, err := clientSet.RbacV1().ClusterRoleBindings().Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1rbac.ClusterRoleBinding{}, err
	}

	return *clusterrolebinding, err
}

func (k *Ingress) Get(opt *metav1.GetOptions) (v1networking.Ingress, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1networking.Ingress{}, err
	}

	ingress, err := clientSet.NetworkingV1().Ingresses(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1networking.Ingress{}, err
	}

	return *ingress, err
}

func (k *HorizontalPodAutoscaler) Get(opt *metav1.GetOptions) (v2autoscaling.HorizontalPodAutoscaler, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v2autoscaling.HorizontalPodAutoscaler{}, err
	}

	hpa, err := clientSet.AutoscalingV2().HorizontalPodAutoscalers(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v2autoscaling.HorizontalPodAutoscaler{}, err
	}

	return *hpa, err
}

func (k *PodDisruptionBudget) Get(opt *metav1.GetOptions) (v1policy.PodDisruptionBudget, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1policy.PodDisruptionBudget{}, err
	}

	pdb, err := clientSet.PolicyV1().PodDisruptionBudgets(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1policy.PodDisruptionBudget{}, err
	}

	return *pdb, err
}

func (k *Secret) Get(opt *metav1.GetOptions) (v1core.Secret, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.Secret{}, err
	}

	secret, err := clientSet.CoreV1().Secrets(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.Secret{}, err
	}

	return *secret, err
}

func (k *Event) Get(opt *metav1.GetOptions) (v1core.Event, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.Event{}, err
	}

	event, err := clientSet.CoreV1().Events(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.Event{}, err
	}

	return *event, err
}

func (k *Pod) Get(opt *metav1.GetOptions) (v1core.Pod, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.Pod{}, err
	}

	pod, err := clientSet.CoreV1().Pods(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.Pod{}, err
	}

	return *pod, err
}

func (k *PersistentVolume) Get(opt *metav1.GetOptions) (v1core.PersistentVolume, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.PersistentVolume{}, err
	}

	pv, err := clientSet.CoreV1().PersistentVolumes().Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.PersistentVolume{}, err
	}

	return *pv, err
}

func (k *Service) Get(opt *metav1.GetOptions) (v1core.Service, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.Service{}, err
	}

	svc, err := clientSet.CoreV1().Services(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.Service{}, err
	}

	return *svc, err
}

func (k *ServiceAccount) Get(opt *metav1.GetOptions) (v1core.ServiceAccount, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.ServiceAccount{}, err
	}

	sa, err := clientSet.CoreV1().ServiceAccounts(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.ServiceAccount{}, err
	}

	return *sa, err
}

func (k *NameSpace) Get(opt *metav1.GetOptions) (v1core.Namespace, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.Namespace{}, err
	}

	sa, err := clientSet.CoreV1().Namespaces().Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.Namespace{}, err
	}

	return *sa, err
}

func (k *Node) Get(opt *metav1.GetOptions) (v1core.Node, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.Node{}, err
	}

	node, err := clientSet.CoreV1().Nodes().Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.Node{}, err
	}

	return *node, err
}

func (k *ConfigMap) Get(opt *metav1.GetOptions) (v1core.ConfigMap, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.ConfigMap{}, err
	}

	cm, err := clientSet.CoreV1().ConfigMaps(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.ConfigMap{}, err
	}

	return *cm, err
}

func (k *PersistentVolumeClaim) Get(opt *metav1.GetOptions) (v1core.PersistentVolumeClaim, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.PersistentVolumeClaim{}, err
	}

	pvc, err := clientSet.CoreV1().PersistentVolumeClaims(k.Namespace).Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1core.PersistentVolumeClaim{}, err
	}

	return *pvc, err
}

func (k *PriorityClass) Get(opt *metav1.GetOptions) (v1scheduling.PriorityClass, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1scheduling.PriorityClass{}, err
	}

	priorityclass, err := clientSet.SchedulingV1().PriorityClasses().Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1scheduling.PriorityClass{}, err
	}

	return *priorityclass, err
}

func (k *StorageClass) Get(opt *metav1.GetOptions) (v1storage.StorageClass, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1storage.StorageClass{}, err
	}

	storageclass, err := clientSet.StorageV1().StorageClasses().Get(context.TODO(), k.Name, *opt)
	if err != nil {
		return v1storage.StorageClass{}, err
	}

	return *storageclass, err
}