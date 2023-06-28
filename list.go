package kube_utils

import (
	"context"
	"errors"
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
	"kube-center/kube-utils/client"
)

func (k *Deployment) GetAll(opt *metav1.ListOptions) (v1app.DeploymentList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1app.DeploymentList{}, err
	}

	deployments, err := clientSet.AppsV1().Deployments(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1app.DeploymentList{}, err
	}

	if len(deployments.Items) == 0 {
		return v1app.DeploymentList{}, errors.New("deployment 不存在")
	}

	return *deployments, nil
}

func (k *DaemonSet) GetAll(opt *metav1.ListOptions) (v1app.DaemonSetList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1app.DaemonSetList{}, err
	}

	daemonsets, err := clientSet.AppsV1().DaemonSets(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1app.DaemonSetList{}, err
	}

	if len(daemonsets.Items) == 0 {
		return v1app.DaemonSetList{}, errors.New("daemonset 不存在")
	}

	return *daemonsets, err
}

func (k *ReplicaSet) GetAll(opt *metav1.ListOptions) (v1app.ReplicaSetList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1app.ReplicaSetList{}, err
	}

	replicasets, err := clientSet.AppsV1().ReplicaSets(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1app.ReplicaSetList{}, err
	}

	if len(replicasets.Items) == 0 {
		return v1app.ReplicaSetList{}, errors.New("replicaset 不存在")
	}

	return *replicasets, err
}

func (k *StatefulSet) GetAll(opt *metav1.ListOptions) (v1app.StatefulSetList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1app.StatefulSetList{}, err
	}

	statefulsets, err := clientSet.AppsV1().StatefulSets(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1app.StatefulSetList{}, err
	}

	if len(statefulsets.Items) == 0 {
		return v1app.StatefulSetList{}, errors.New("statefulset 不存在")
	}

	return *statefulsets, err
}

func (k *CronJob) GetAll(opt *metav1.ListOptions) (v1batch.CronJobList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1batch.CronJobList{}, err
	}

	cronjobs, err := clientSet.BatchV1().CronJobs(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1batch.CronJobList{}, err
	}

	if len(cronjobs.Items) == 0 {
		return v1batch.CronJobList{}, errors.New("cronjob 不存在")
	}

	return *cronjobs, err
}

func (k *Job) GetAll(opt *metav1.ListOptions) (v1batch.JobList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1batch.JobList{}, err
	}

	jobs, err := clientSet.BatchV1().Jobs(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1batch.JobList{}, err
	}

	if len(jobs.Items) == 0 {
		return v1batch.JobList{}, errors.New("job 不存在")
	}

	return *jobs, err
}

func (k *Role) GetAll(opt *metav1.ListOptions) (v1rbac.RoleList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1rbac.RoleList{}, err
	}

	roles, err := clientSet.RbacV1().Roles(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1rbac.RoleList{}, err
	}

	if len(roles.Items) == 0 {
		return v1rbac.RoleList{}, errors.New("role 不存在")
	}

	return *roles, err
}

func (k *RoleBinding) GetAll(opt *metav1.ListOptions) (v1rbac.RoleBindingList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1rbac.RoleBindingList{}, err
	}

	rolebindings, err := clientSet.RbacV1().RoleBindings(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1rbac.RoleBindingList{}, err
	}

	if len(rolebindings.Items) == 0 {
		return v1rbac.RoleBindingList{}, errors.New("rolebinding 不存在")
	}

	return *rolebindings, err
}

func (k *ClusterRole) GetAll(opt *metav1.ListOptions) (v1rbac.ClusterRoleList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1rbac.ClusterRoleList{}, err
	}

	clusterroles, err := clientSet.RbacV1().ClusterRoles().List(context.TODO(), *opt)
	if err != nil {
		return v1rbac.ClusterRoleList{}, err
	}

	if len(clusterroles.Items) == 0 {
		return v1rbac.ClusterRoleList{}, errors.New("clusterrole 不存在")
	}

	return *clusterroles, err
}

func (k *ClusterRoleBinding) GetAll(opt *metav1.ListOptions) (v1rbac.ClusterRoleBindingList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1rbac.ClusterRoleBindingList{}, err
	}

	clusterrolebindings, err := clientSet.RbacV1().ClusterRoleBindings().List(context.TODO(), *opt)
	if err != nil {
		return v1rbac.ClusterRoleBindingList{}, err
	}

	if len(clusterrolebindings.Items) == 0 {
		return v1rbac.ClusterRoleBindingList{}, errors.New("clusterrolebinding 不存在")
	}

	return *clusterrolebindings, err
}

func (k *Ingress) GetAll(opt *metav1.ListOptions) (v1networking.IngressList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1networking.IngressList{}, err
	}

	ingresses, err := clientSet.NetworkingV1().Ingresses(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1networking.IngressList{}, err
	}

	if len(ingresses.Items) == 0 {
		return v1networking.IngressList{}, errors.New("ingress 不存在")
	}

	return *ingresses, err
}

func (k *HorizontalPodAutoscaler) GetAll(opt *metav1.ListOptions) (v2autoscaling.HorizontalPodAutoscalerList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v2autoscaling.HorizontalPodAutoscalerList{}, err
	}

	hpas, err := clientSet.AutoscalingV2().HorizontalPodAutoscalers(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v2autoscaling.HorizontalPodAutoscalerList{}, err
	}

	if len(hpas.Items) == 0 {
		return v2autoscaling.HorizontalPodAutoscalerList{}, errors.New("hpa 不存在")
	}

	return *hpas, err
}

func (k *PodDisruptionBudget) GetAll(opt *metav1.ListOptions) (v1policy.PodDisruptionBudgetList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1policy.PodDisruptionBudgetList{}, err
	}

	pdbs, err := clientSet.PolicyV1().PodDisruptionBudgets(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1policy.PodDisruptionBudgetList{}, err
	}

	if len(pdbs.Items) == 0 {
		return v1policy.PodDisruptionBudgetList{}, errors.New("pdb 不存在")
	}

	return *pdbs, err
}

func (k *Secret) GetAll(opt *metav1.ListOptions) (v1core.SecretList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.SecretList{}, err
	}

	secrets, err := clientSet.CoreV1().Secrets(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1core.SecretList{}, err
	}

	if len(secrets.Items) == 0 {
		return v1core.SecretList{}, errors.New("secret 不存在")
	}

	return *secrets, err
}

func (k *Event) GetAll(opt *metav1.ListOptions) (v1core.EventList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.EventList{}, err
	}

	events, err := clientSet.CoreV1().Events(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1core.EventList{}, err
	}

	if len(events.Items) == 0 {
		return v1core.EventList{}, errors.New("event 不存在")
	}

	return *events, err
}

func (k *Pod) GetAll(opt *metav1.ListOptions) (v1core.PodList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.PodList{}, err
	}

	pods, err := clientSet.CoreV1().Pods(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1core.PodList{}, err
	}

	if len(pods.Items) == 0 {
		return v1core.PodList{}, errors.New("pod 不存在")
	}

	return *pods, err
}

func (k *PersistentVolume) GetAll(opt *metav1.ListOptions) (v1core.PersistentVolumeList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.PersistentVolumeList{}, err
	}

	pvs, err := clientSet.CoreV1().PersistentVolumes().List(context.TODO(), *opt)
	if err != nil {
		return v1core.PersistentVolumeList{}, err
	}

	if len(pvs.Items) == 0 {
		return v1core.PersistentVolumeList{}, errors.New("pv 不存在")
	}

	return *pvs, err
}

func (k *Service) GetAll(opt *metav1.ListOptions) (v1core.ServiceList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.ServiceList{}, err
	}

	svcs, err := clientSet.CoreV1().Services(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1core.ServiceList{}, err
	}

	if len(svcs.Items) == 0 {
		return v1core.ServiceList{}, errors.New("service 不存在")
	}

	return *svcs, err
}

func (k *ServiceAccount) GetAll(opt *metav1.ListOptions) (v1core.ServiceAccountList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.ServiceAccountList{}, err
	}

	sas, err := clientSet.CoreV1().ServiceAccounts(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1core.ServiceAccountList{}, err
	}

	if len(sas.Items) == 0 {
		return v1core.ServiceAccountList{}, errors.New("sa 不存在")
	}

	return *sas, err
}

func (k *NameSpace) GetAll(opt *metav1.ListOptions) (v1core.NamespaceList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.NamespaceList{}, err
	}

	namespaces, err := clientSet.CoreV1().Namespaces().List(context.TODO(), *opt)
	if err != nil {
		return v1core.NamespaceList{}, err
	}

	if len(namespaces.Items) == 0 {
		return v1core.NamespaceList{}, errors.New("namespace 不存在")
	}

	return *namespaces, err
}

func (k *Node) GetAll(opt *metav1.ListOptions) (v1core.NodeList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.NodeList{}, err
	}

	nodes, err := clientSet.CoreV1().Nodes().List(context.TODO(), *opt)
	if err != nil {
		return v1core.NodeList{}, err
	}

	if len(nodes.Items) == 0 {
		return v1core.NodeList{}, errors.New("node 不存在")
	}

	return *nodes, err
}

func (k *ConfigMap) GetAll(opt *metav1.ListOptions) (v1core.ConfigMapList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.ConfigMapList{}, err
	}

	cms, err := clientSet.CoreV1().ConfigMaps(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1core.ConfigMapList{}, err
	}

	if len(cms.Items) == 0 {
		return v1core.ConfigMapList{}, errors.New("configmap 不存在")
	}

	return *cms, err
}

func (k *PersistentVolumeClaim) GetAll(opt *metav1.ListOptions) (v1core.PersistentVolumeClaimList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1core.PersistentVolumeClaimList{}, err
	}

	pvcs, err := clientSet.CoreV1().PersistentVolumeClaims(k.Namespace).List(context.TODO(), *opt)
	if err != nil {
		return v1core.PersistentVolumeClaimList{}, err
	}

	if len(pvcs.Items) == 0 {
		return v1core.PersistentVolumeClaimList{}, errors.New("pvc 不存在")
	}

	return *pvcs, err
}

func (k *PriorityClass) GetAll(opt *metav1.ListOptions) (v1scheduling.PriorityClassList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1scheduling.PriorityClassList{}, err
	}

	priorityclasses, err := clientSet.SchedulingV1().PriorityClasses().List(context.TODO(), *opt)
	if err != nil {
		return v1scheduling.PriorityClassList{}, err
	}

	if len(priorityclasses.Items) == 0 {
		return v1scheduling.PriorityClassList{}, errors.New("PriorityClass 不存在")
	}

	return *priorityclasses, err
}

func (k *StorageClass) GetAll(opt *metav1.ListOptions) (v1storage.StorageClassList, error) {
	var clientSet, err = client.InitClient()
	if err != nil {
		return v1storage.StorageClassList{}, err
	}

	storageclasses, err := clientSet.StorageV1().StorageClasses().List(context.TODO(), *opt)
	if err != nil {
		return v1storage.StorageClassList{}, err
	}

	if len(storageclasses.Items) == 0 {
		return v1storage.StorageClassList{}, errors.New("StorageClass 不存在")
	}

	return *storageclasses, err
}
