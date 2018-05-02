package v1

import (
	"fmt"

	"github.com/appscode/kubernetes-webhook-util/apis/workload/v1"
	ocapps "github.com/openshift/api/apps/v1"
	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	core "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

func NewWorkload(t metav1.TypeMeta, o metav1.ObjectMeta, tpl core.PodTemplateSpec) *v1.Workload {
	return &v1.Workload{
		TypeMeta:   t,
		ObjectMeta: o,
		Spec: v1.WorkloadSpec{
			Template: tpl,
		},
	}
}

func NewObjectForGVK(gvk schema.GroupVersionKind, name, ns string) (runtime.Object, error) {
	obj, err := legacyscheme.Scheme.New(gvk)
	if err != nil {
		return nil, err
	}
	obj.GetObjectKind().SetGroupVersionKind(gvk)
	out, err := meta.Accessor(obj)
	if err != nil {
		return nil, err
	}
	out.SetName(name)
	out.SetNamespace(ns)
	return obj, nil
}

func NewObject(kindOrResource string, name, ns string) (runtime.Object, error) {
	switch kindOrResource {
	case v1.KindPod, v1.ResourcePods, v1.ResourcePod:
		return &core.Pod{
			TypeMeta:   metav1.TypeMeta{APIVersion: core.SchemeGroupVersion.String(), Kind: v1.KindPod},
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		}, nil
	case v1.KindReplicationController, v1.ResourceReplicationControllers, v1.ResourceReplicationController:
		return &core.ReplicationController{
			TypeMeta:   metav1.TypeMeta{APIVersion: core.SchemeGroupVersion.String(), Kind: v1.KindReplicationController},
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		}, nil
	case v1.KindDeployment, v1.ResourceDeployments, v1.ResourceDeployment:
		return &appsv1beta1.Deployment{
			TypeMeta:   metav1.TypeMeta{APIVersion: appsv1beta1.SchemeGroupVersion.String(), Kind: v1.KindDeployment},
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		}, nil
	case v1.KindDaemonSet, v1.ResourceDaemonSets, v1.ResourceDaemonSet:
		return &extensions.DaemonSet{
			TypeMeta:   metav1.TypeMeta{APIVersion: extensions.SchemeGroupVersion.String(), Kind: v1.KindDaemonSet},
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		}, nil
	case v1.KindReplicaSet, v1.ResourceReplicaSets, v1.ResourceReplicaSet:
		return &extensions.ReplicaSet{
			TypeMeta:   metav1.TypeMeta{APIVersion: extensions.SchemeGroupVersion.String(), Kind: v1.KindReplicaSet},
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		}, nil
	case v1.KindStatefulSet, v1.ResourceStatefulSets, v1.ResourceStatefulSet:
		return &appsv1beta1.StatefulSet{
			TypeMeta:   metav1.TypeMeta{APIVersion: appsv1beta1.SchemeGroupVersion.String(), Kind: v1.KindStatefulSet},
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		}, nil
	case v1.KindJob, v1.ResourceJobs, v1.ResourceJob:
		return &batchv1.Job{
			TypeMeta:   metav1.TypeMeta{APIVersion: batchv1.SchemeGroupVersion.String(), Kind: v1.KindJob},
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		}, nil
	case v1.KindCronJob, v1.ResourceCronJobs, v1.ResourceCronJob:
		return &batchv1beta1.CronJob{
			TypeMeta:   metav1.TypeMeta{APIVersion: batchv1beta1.SchemeGroupVersion.String(), Kind: v1.KindCronJob},
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		}, nil
	case v1.KindDeploymentConfig, v1.ResourceDeploymentConfigs, v1.ResourceDeploymentConfig:
		return &ocapps.DeploymentConfig{
			TypeMeta:   metav1.TypeMeta{APIVersion: ocapps.SchemeGroupVersion.String(), Kind: v1.KindDeploymentConfig},
			ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: ns},
		}, nil
	default:
		return nil, fmt.Errorf("unknown kind or resource %s", kindOrResource)
	}
}

func newWithObject(t metav1.TypeMeta, o metav1.ObjectMeta, replicas *int32, sel *metav1.LabelSelector, tpl core.PodTemplateSpec, obj runtime.Object) *v1.Workload {
	return &v1.Workload{
		TypeMeta:   t,
		ObjectMeta: o,
		Spec: v1.WorkloadSpec{
			Replicas: replicas,
			Selector: sel,
			Template: tpl,
		},
		Object: obj,
	}
}

// ref: https://github.com/kubernetes/kubernetes/blob/4f083dee54539b0ca24ddc55d53921f5c2efc0b9/pkg/kubectl/cmd/util/factory_client_access.go#L221
func ConvertToWorkload(obj runtime.Object) (*v1.Workload, error) {
	switch t := obj.(type) {
	case *core.Pod:
		return newWithObject(t.TypeMeta, t.ObjectMeta, nil, nil, core.PodTemplateSpec{ObjectMeta: t.ObjectMeta, Spec: t.Spec}, obj), nil
		// ReplicationController
	case *core.ReplicationController:
		if t.Spec.Template == nil {
			t.Spec.Template = &core.PodTemplateSpec{}
		}
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, &metav1.LabelSelector{MatchLabels: t.Spec.Selector}, *t.Spec.Template, obj), nil
		// Deployment
	case *extensions.Deployment:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
	case *appsv1beta1.Deployment:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
	case *appsv1beta2.Deployment:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
	case *appsv1.Deployment:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
		// DaemonSet
	case *extensions.DaemonSet:
		return newWithObject(t.TypeMeta, t.ObjectMeta, nil, t.Spec.Selector, t.Spec.Template, obj), nil
	case *appsv1beta2.DaemonSet:
		return newWithObject(t.TypeMeta, t.ObjectMeta, nil, t.Spec.Selector, t.Spec.Template, obj), nil
	case *appsv1.DaemonSet:
		return newWithObject(t.TypeMeta, t.ObjectMeta, nil, t.Spec.Selector, t.Spec.Template, obj), nil
		// ReplicaSet
	case *extensions.ReplicaSet:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
	case *appsv1beta2.ReplicaSet:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
	case *appsv1.ReplicaSet:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
		// StatefulSet
	case *appsv1beta1.StatefulSet:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
	case *appsv1beta2.StatefulSet:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
	case *appsv1.StatefulSet:
		return newWithObject(t.TypeMeta, t.ObjectMeta, t.Spec.Replicas, t.Spec.Selector, t.Spec.Template, obj), nil
		// Job
	case *batchv1.Job:
		return newWithObject(t.TypeMeta, t.ObjectMeta, nil, t.Spec.Selector, t.Spec.Template, obj), nil
		// CronJob
	case *batchv1beta1.CronJob:
		return newWithObject(t.TypeMeta, t.ObjectMeta, nil, t.Spec.JobTemplate.Spec.Selector, t.Spec.JobTemplate.Spec.Template, obj), nil
		// DeploymentConfig
	case *ocapps.DeploymentConfig:
		if t.Spec.Template == nil {
			t.Spec.Template = &core.PodTemplateSpec{}
		}
		var replicas *int32
		if t.Spec.Replicas > 0 {
			replicas = &t.Spec.Replicas
		}
		return newWithObject(t.TypeMeta, t.ObjectMeta, replicas, &metav1.LabelSelector{MatchLabels: t.Spec.Selector}, *t.Spec.Template, obj), nil
	case *v1.Workload:
		return t, nil
	default:
		return nil, fmt.Errorf("the object is not a pod or does not have a pod template")
	}
}

func ApplyWorkload(obj runtime.Object, w *v1.Workload) error {
	switch t := obj.(type) {
	case *core.Pod:
		t.ObjectMeta = w.ObjectMeta
		t.Spec = w.Spec.Template.Spec
		// ReplicationController
	case *core.ReplicationController:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = &w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
		// Deployment
	case *extensions.Deployment:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
	case *appsv1beta1.Deployment:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
	case *appsv1beta2.Deployment:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
	case *appsv1.Deployment:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
		// DaemonSet
	case *extensions.DaemonSet:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
	case *appsv1beta2.DaemonSet:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
	case *appsv1.DaemonSet:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		// ReplicaSet
	case *extensions.ReplicaSet:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
	case *appsv1beta2.ReplicaSet:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
	case *appsv1.ReplicaSet:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
		// StatefulSet
	case *appsv1beta1.StatefulSet:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
	case *appsv1beta2.StatefulSet:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
	case *appsv1.StatefulSet:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = w.Spec.Replicas
		}
		// Job
	case *batchv1.Job:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = w.Spec.Template
		// CronJob
	case *batchv1beta1.CronJob:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.JobTemplate.Spec.Template = w.Spec.Template
		// DeploymentConfig
	case *ocapps.DeploymentConfig:
		t.ObjectMeta = w.ObjectMeta
		t.Spec.Template = &w.Spec.Template
		if w.Spec.Replicas != nil {
			t.Spec.Replicas = *w.Spec.Replicas
		}
	default:
		return fmt.Errorf("the object is not a pod or does not have a pod template")
	}
	return nil
}
