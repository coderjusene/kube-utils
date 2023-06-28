package kube_utils

import (
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
	"k8s.io/apimachinery/pkg/types"
)

type DoDeployment interface {
	Get(opt *metav1.GetOptions) (v1app.Deployment, error)
	GetAll(opt *metav1.ListOptions) (v1app.DeploymentList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Delete(opt *metav1.DeleteOptions) error
	Patch(data string, pt *types.PatchType, opt *metav1.PatchOptions) (v1app.Deployment, error)
	Update(yaml []byte, opt *metav1.UpdateOptions) error
	Scale(replicas int32) error
	UpdateImage(tag string) error
}

type DoDaemonSet interface {
	Get(opt *metav1.GetOptions) (v1app.DaemonSet, error)
	GetAll(opt *metav1.ListOptions) (v1app.DaemonSetList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Delete(opt *metav1.DeleteOptions) error
	Update(yaml []byte, opt *metav1.UpdateOptions) error
	Patch(data string, pt *types.PatchType, opt *metav1.PatchOptions) (v1app.DaemonSet, error)
	UpdateImage(tag string) error
}

type DoReplicaSet interface {
	Get(opt *metav1.GetOptions) (v1app.ReplicaSet, error)
	GetAll(opt *metav1.ListOptions) (v1app.ReplicaSetList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Delete(opt *metav1.DeleteOptions) error
	Update(yaml []byte, opt *metav1.UpdateOptions) error
}

type DoStatefulSet interface {
	Get(opt *metav1.GetOptions) (v1app.StatefulSet, error)
	GetAll(opt *metav1.ListOptions) (v1app.StatefulSetList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Delete(opt *metav1.DeleteOptions) error
	Patch(data string, pt *types.PatchType, opt *metav1.PatchOptions) (v1app.StatefulSet, error)
	Scale(replicas int32) error
	UpdateImage(tag string) error
}

type DoCronJob interface {
	Get(opt *metav1.GetOptions) (v1batch.CronJob, error)
	GetAll(opt *metav1.ListOptions) (v1batch.CronJobList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Delete(opt *metav1.DeleteOptions) error
}

type DoJob interface {
	Get(opt *metav1.GetOptions) (v1batch.Job, error)
	GetAll(opt *metav1.ListOptions) (v1batch.JobList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Delete(opt *metav1.DeleteOptions) error
}

type DoRole interface {
	Get(opt *metav1.GetOptions) (v1rbac.Role, error)
	GetAll(opt *metav1.ListOptions) (v1rbac.RoleList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoRoleBinding interface {
	Get(opt *metav1.GetOptions) (v1rbac.RoleBinding, error)
	GetAll(opt *metav1.ListOptions) (v1rbac.RoleBindingList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoClusterRole interface {
	Get(opt *metav1.GetOptions) (v1rbac.ClusterRole, error)
	GetAll(opt *metav1.ListOptions) (v1rbac.ClusterRoleList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoClusterRoleBinding interface {
	Get(opt *metav1.GetOptions) (v1rbac.ClusterRoleBinding, error)
	GetAll(opt *metav1.ListOptions) (v1rbac.ClusterRoleBindingList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoIngress interface {
	Get(opt *metav1.GetOptions) (v1networking.Ingress, error)
	GetAll(opt *metav1.ListOptions) (v1networking.IngressList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Update(yaml []byte, opt *metav1.UpdateOptions) error
	Delete(opt *metav1.DeleteOptions) error
}

type DoHorizontalPodAutoscaler interface {
	Get(opt *metav1.GetOptions) (v2autoscaling.HorizontalPodAutoscaler, error)
	GetAll(opt *metav1.ListOptions) (v2autoscaling.HorizontalPodAutoscalerList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoPodDisruptionBudget interface {
	Get(opt *metav1.GetOptions) (v1policy.PodDisruptionBudget, error)
	GetAll(opt *metav1.ListOptions) (v1policy.PodDisruptionBudgetList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoSecret interface {
	Get(opt *metav1.GetOptions) (v1core.Secret, error)
	GetAll(opt *metav1.ListOptions) (v1core.SecretList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoEvent interface {
	Get(opt *metav1.GetOptions) (v1core.Event, error)
	GetAll(opt *metav1.ListOptions) (v1core.EventList, error)
}

type DoPod interface {
	Get(opt *metav1.GetOptions) (v1core.Pod, error)
	GetAll(opt *metav1.ListOptions) (v1core.PodList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Logs(podLogOpts *v1core.PodLogOptions) (string, error)
	Exec(cmd []string) error
	CopyToPod(src, dest string) error
	CopyFromPod(src, dest string) error
	Describe() (v1core.Pod, v1core.EventList, error)
	Delete(opt *metav1.DeleteOptions) error
}

type DoPersistentVolume interface {
	Get(opt *metav1.GetOptions) (v1core.PersistentVolume, error)
	GetAll(opt *metav1.ListOptions) (v1core.PersistentVolumeList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoService interface {
	Get(opt *metav1.GetOptions) (v1core.Service, error)
	GetAll(opt *metav1.ListOptions) (v1core.ServiceList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Delete(opt *metav1.DeleteOptions) error
	Update(yaml []byte, opt *metav1.UpdateOptions) error
}

type DoServiceAccount interface {
	Get(opt *metav1.GetOptions) (v1core.ServiceAccount, error)
	GetAll(opt *metav1.ListOptions) (v1core.ServiceAccountList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoNameSpace interface {
	Get(opt *metav1.GetOptions) (v1core.Namespace, error)
	GetAll(opt *metav1.ListOptions) (v1core.NamespaceList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
	Delete(opt *metav1.DeleteOptions) error
}

type DoNode interface {
	Get(opt *metav1.GetOptions) (v1core.Node, error)
	GetAll(opt *metav1.ListOptions) (v1core.NodeList, error)
	Describe() (v1core.Node, v1core.EventList, error)
}

type DoConfigMap interface {
	Get(opt *metav1.GetOptions) (v1core.ConfigMap, error)
	GetAll(opt *metav1.ListOptions) (v1core.ConfigMapList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoPersistentVolumeClaim interface {
	Get(opt *metav1.GetOptions) (v1core.PersistentVolumeClaim, error)
	GetAll(opt *metav1.ListOptions) (v1core.PersistentVolumeClaimList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoPriorityClass interface {
	Get(opt *metav1.GetOptions) (v1scheduling.PriorityClass, error)
	GetAll(opt *metav1.ListOptions) (v1scheduling.PriorityClassList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}

type DoStorageClass interface {
	Get(opt *metav1.GetOptions) (v1storage.StorageClass, error)
	GetAll(opt *metav1.ListOptions) (v1storage.StorageClassList, error)
	Create(yaml []byte, opt *metav1.CreateOptions) error
}
