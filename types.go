package kube_utils

// apps.v1
type DaemonSet struct {
	Name      string
	Namespace string
}

type Deployment struct {
	Name      string
	Namespace string
}

type ReplicaSet struct {
	Name      string
	Namespace string
}

type StatefulSet struct {
	Name      string
	Namespace string
}

// batch/v1
type CronJob struct {
	Name      string
	Namespace string
}
type Job struct {
	Name      string
	Namespace string
}

// rbac.v1
type RoleBinding struct {
	Name      string
	Namespace string
}

type Role struct {
	Name      string
	Namespace string
}

type ClusterRoleBinding struct {
	Name string
}

type ClusterRole struct {
	Name string
}

type Ingress struct {
	Name      string
	Namespace string
}

type HorizontalPodAutoscaler struct {
	Name      string
	Namespace string
}

type PodDisruptionBudget struct {
	Name      string
	Namespace string
}

type Secret struct {
	Name      string
	Namespace string
}

type Event struct {
	Name      string
	Namespace string
}

type Pod struct {
	Name          string
	Namespace     string
	ContainerName string
}

type PersistentVolume struct {
	Name string
}

type PersistentVolumeClaim struct {
	Name      string
	Namespace string
}

type Service struct {
	Name      string
	Namespace string
}

type ServiceAccount struct {
	Name      string
	Namespace string
}

type NameSpace struct {
	Name string
}

type Node struct {
	Name string
}

type ConfigMap struct {
	Name      string
	Namespace string
}

type PriorityClass struct {
	Name string
}

type StorageClass struct {
	Name string
}
