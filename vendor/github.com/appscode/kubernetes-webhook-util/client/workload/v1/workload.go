package v1

import (
	"fmt"

	"github.com/appscode/kubernetes-webhook-util/apis/workload/v1"
	"github.com/appscode/kutil"
	"github.com/evanphx/json-patch"
	"github.com/golang/glog"
	"github.com/json-iterator/go"
	ocapps "github.com/openshift/api/apps/v1"
	occ "github.com/openshift/client-go/apps/clientset/versioned"
	appsv1 "k8s.io/api/apps/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	core "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

var json = jsoniter.ConfigFastest

type WorkloadTransformerFunc func(*v1.Workload) *v1.Workload

// WorkloadsGetter has a method to return a WorkloadInterface.
// A group's client should implement this interface.
type WorkloadsGetter interface {
	Workloads(namespace string) WorkloadInterface
}

// WorkloadInterface has methods to work with Workload resources.
type WorkloadInterface interface {
	Create(*v1.Workload) (*v1.Workload, error)
	Update(*v1.Workload) (*v1.Workload, error)
	Delete(obj runtime.Object, options *metav1.DeleteOptions) error
	Get(obj runtime.Object, options metav1.GetOptions) (*v1.Workload, error)
	List(opts metav1.ListOptions) (*v1.WorkloadList, error)
	Patch(cur *v1.Workload, transform WorkloadTransformerFunc) (*v1.Workload, kutil.VerbType, error)
	PatchObject(cur, mod *v1.Workload) (*v1.Workload, kutil.VerbType, error)
	CreateOrPatch(obj runtime.Object, transform WorkloadTransformerFunc) (*v1.Workload, kutil.VerbType, error)
}

// workloads implements WorkloadInterface
type workloads struct {
	kc kubernetes.Interface
	oc occ.Interface
	ns string
}

var _ WorkloadInterface = &workloads{}

// newWorkloads returns a Workloads
func newWorkloads(kc kubernetes.Interface, oc occ.Interface, namespace string) *workloads {
	return &workloads{
		kc: kc,
		oc: oc,
		ns: namespace,
	}
}

func (c *workloads) Create(w *v1.Workload) (*v1.Workload, error) {
	var out runtime.Object
	var err error
	switch w.GroupVersionKind() {
	case core.SchemeGroupVersion.WithKind("Pod"):
		obj := &core.Pod{}
		ApplyWorkload(obj, w)
		out, err = c.kc.CoreV1().Pods(c.ns).Create(obj)
		// ReplicationController
	case core.SchemeGroupVersion.WithKind("ReplicationController"):
		obj := &core.ReplicationController{}
		ApplyWorkload(obj, w)
		out, err = c.kc.CoreV1().ReplicationControllers(c.ns).Create(obj)
		// Deployment
	case extensions.SchemeGroupVersion.WithKind("Deployment"):
		obj := &extensions.Deployment{}
		ApplyWorkload(obj, w)
		out, err = c.kc.ExtensionsV1beta1().Deployments(c.ns).Create(obj)
	case appsv1beta1.SchemeGroupVersion.WithKind("Deployment"):
		obj := &appsv1beta1.Deployment{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1beta1().Deployments(c.ns).Create(obj)
	case appsv1beta2.SchemeGroupVersion.WithKind("Deployment"):
		obj := &appsv1beta2.Deployment{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1beta2().Deployments(c.ns).Create(obj)
	case appsv1.SchemeGroupVersion.WithKind("Deployment"):
		obj := &appsv1.Deployment{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1().Deployments(c.ns).Create(obj)
		// DaemonSet
	case extensions.SchemeGroupVersion.WithKind("DaemonSet"):
		obj := &extensions.DaemonSet{}
		ApplyWorkload(obj, w)
		out, err = c.kc.ExtensionsV1beta1().DaemonSets(c.ns).Create(obj)
	case appsv1beta2.SchemeGroupVersion.WithKind("DaemonSet"):
		obj := &appsv1beta2.DaemonSet{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1beta2().DaemonSets(c.ns).Create(obj)
	case appsv1.SchemeGroupVersion.WithKind("DaemonSet"):
		obj := &appsv1.DaemonSet{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1().DaemonSets(c.ns).Create(obj)
		// ReplicaSet
	case extensions.SchemeGroupVersion.WithKind("ReplicaSet"):
		obj := &extensions.ReplicaSet{}
		ApplyWorkload(obj, w)
		out, err = c.kc.ExtensionsV1beta1().ReplicaSets(c.ns).Create(obj)
	case appsv1beta2.SchemeGroupVersion.WithKind("ReplicaSet"):
		obj := &appsv1beta2.ReplicaSet{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1beta2().ReplicaSets(c.ns).Create(obj)
	case appsv1.SchemeGroupVersion.WithKind("ReplicaSet"):
		obj := &appsv1.ReplicaSet{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1().ReplicaSets(c.ns).Create(obj)
		// StatefulSet
	case appsv1beta1.SchemeGroupVersion.WithKind("StatefulSet"):
		obj := &appsv1beta1.StatefulSet{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1beta1().StatefulSets(c.ns).Create(obj)
	case appsv1beta2.SchemeGroupVersion.WithKind("StatefulSet"):
		obj := &appsv1beta2.StatefulSet{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1beta2().StatefulSets(c.ns).Create(obj)
	case appsv1.SchemeGroupVersion.WithKind("StatefulSet"):
		obj := &appsv1.StatefulSet{}
		ApplyWorkload(obj, w)
		out, err = c.kc.AppsV1().StatefulSets(c.ns).Create(obj)
		// Job
	case batchv1.SchemeGroupVersion.WithKind("Job"):
		obj := &batchv1.Job{}
		ApplyWorkload(obj, w)
		out, err = c.kc.BatchV1().Jobs(c.ns).Create(obj)
		// CronJob
	case batchv1beta1.SchemeGroupVersion.WithKind("CronJob"):
		obj := &batchv1beta1.CronJob{}
		ApplyWorkload(obj, w)
		out, err = c.kc.BatchV1beta1().CronJobs(c.ns).Create(obj)
	case ocapps.SchemeGroupVersion.WithKind("DeploymentConfig"):
		obj := &ocapps.DeploymentConfig{}
		ApplyWorkload(obj, w)
		out, err = c.oc.AppsV1().DeploymentConfigs(c.ns).Create(obj)
	default:
		err = fmt.Errorf("the object is not a pod or does not have a pod template")
	}
	if err != nil {
		return nil, err
	}
	return ConvertToWorkload(out)
}

func (c *workloads) Update(w *v1.Workload) (*v1.Workload, error) {
	var out runtime.Object
	var err error
	switch t := w.Object.(type) {
	case *core.Pod:
		out, err = c.kc.CoreV1().Pods(c.ns).Update(t)
		// ReplicationController
	case *core.ReplicationController:
		out, err = c.kc.CoreV1().ReplicationControllers(c.ns).Update(t)
		// Deployment
	case *extensions.Deployment:
		out, err = c.kc.ExtensionsV1beta1().Deployments(c.ns).Update(t)
	case *appsv1beta1.Deployment:
		out, err = c.kc.AppsV1beta1().Deployments(c.ns).Update(t)
	case *appsv1beta2.Deployment:
		out, err = c.kc.AppsV1beta2().Deployments(c.ns).Update(t)
	case *appsv1.Deployment:
		out, err = c.kc.AppsV1().Deployments(c.ns).Update(t)
		// DaemonSet
	case *extensions.DaemonSet:
		out, err = c.kc.ExtensionsV1beta1().DaemonSets(c.ns).Update(t)
	case *appsv1beta2.DaemonSet:
		out, err = c.kc.AppsV1beta2().DaemonSets(c.ns).Update(t)
	case *appsv1.DaemonSet:
		out, err = c.kc.AppsV1().DaemonSets(c.ns).Update(t)
		// ReplicaSet
	case *extensions.ReplicaSet:
		out, err = c.kc.ExtensionsV1beta1().ReplicaSets(c.ns).Update(t)
	case *appsv1beta2.ReplicaSet:
		out, err = c.kc.AppsV1beta2().ReplicaSets(c.ns).Update(t)
	case *appsv1.ReplicaSet:
		out, err = c.kc.AppsV1().ReplicaSets(c.ns).Update(t)
		// StatefulSet
	case *appsv1beta1.StatefulSet:
		out, err = c.kc.AppsV1beta1().StatefulSets(c.ns).Update(t)
	case *appsv1beta2.StatefulSet:
		out, err = c.kc.AppsV1beta2().StatefulSets(c.ns).Update(t)
	case *appsv1.StatefulSet:
		out, err = c.kc.AppsV1().StatefulSets(c.ns).Update(t)
		// Job
	case *batchv1.Job:
		out, err = c.kc.BatchV1().Jobs(c.ns).Update(t)
		// CronJob
	case *batchv1beta1.CronJob:
		out, err = c.kc.BatchV1beta1().CronJobs(c.ns).Update(t)
	case *ocapps.DeploymentConfig:
		out, err = c.oc.AppsV1().DeploymentConfigs(c.ns).Update(t)
	default:
		err = fmt.Errorf("the object is not a pod or does not have a pod template")
	}
	if err != nil {
		return nil, err
	}
	return ConvertToWorkload(out)
}

func (c *workloads) Delete(obj runtime.Object, options *metav1.DeleteOptions) error {
	switch t := obj.(type) {
	case *core.Pod:
		return c.kc.CoreV1().Pods(c.ns).Delete(t.ObjectMeta.Name, options)
		// ReplicationController
	case *core.ReplicationController:
		return c.kc.CoreV1().ReplicationControllers(c.ns).Delete(t.ObjectMeta.Name, options)
		// Deployment
	case *extensions.Deployment:
		return c.kc.ExtensionsV1beta1().Deployments(c.ns).Delete(t.ObjectMeta.Name, options)
	case *appsv1beta1.Deployment:
		return c.kc.AppsV1beta1().Deployments(c.ns).Delete(t.ObjectMeta.Name, options)
	case *appsv1beta2.Deployment:
		return c.kc.AppsV1beta2().Deployments(c.ns).Delete(t.ObjectMeta.Name, options)
	case *appsv1.Deployment:
		return c.kc.AppsV1().Deployments(c.ns).Delete(t.ObjectMeta.Name, options)
		// DaemonSet
	case *extensions.DaemonSet:
		return c.kc.ExtensionsV1beta1().DaemonSets(c.ns).Delete(t.ObjectMeta.Name, options)
	case *appsv1beta2.DaemonSet:
		return c.kc.AppsV1beta2().DaemonSets(c.ns).Delete(t.ObjectMeta.Name, options)
	case *appsv1.DaemonSet:
		return c.kc.AppsV1().DaemonSets(c.ns).Delete(t.ObjectMeta.Name, options)
		// ReplicaSet
	case *extensions.ReplicaSet:
		return c.kc.ExtensionsV1beta1().ReplicaSets(c.ns).Delete(t.ObjectMeta.Name, options)
	case *appsv1beta2.ReplicaSet:
		return c.kc.AppsV1beta2().ReplicaSets(c.ns).Delete(t.ObjectMeta.Name, options)
	case *appsv1.ReplicaSet:
		return c.kc.AppsV1().ReplicaSets(c.ns).Delete(t.ObjectMeta.Name, options)
		// StatefulSet
	case *appsv1beta1.StatefulSet:
		return c.kc.AppsV1beta1().StatefulSets(c.ns).Delete(t.ObjectMeta.Name, options)
	case *appsv1beta2.StatefulSet:
		return c.kc.AppsV1beta2().StatefulSets(c.ns).Delete(t.ObjectMeta.Name, options)
	case *appsv1.StatefulSet:
		return c.kc.AppsV1().StatefulSets(c.ns).Delete(t.ObjectMeta.Name, options)
		// Job
	case *batchv1.Job:
		return c.kc.BatchV1().Jobs(c.ns).Delete(t.ObjectMeta.Name, options)
		// CronJob
	case *batchv1beta1.CronJob:
		return c.kc.BatchV1beta1().CronJobs(c.ns).Delete(t.ObjectMeta.Name, options)
	case *ocapps.DeploymentConfig:
		return c.oc.AppsV1().DeploymentConfigs(c.ns).Delete(t.ObjectMeta.Name, options)
	default:
		return fmt.Errorf("the object is not a pod or does not have a pod template")
	}
}

func (c *workloads) Get(obj runtime.Object, options metav1.GetOptions) (*v1.Workload, error) {
	var out runtime.Object
	var err error
	switch t := obj.(type) {
	case *core.Pod:
		out, err = c.kc.CoreV1().Pods(c.ns).Get(t.ObjectMeta.Name, options)
		// ReplicationController
	case *core.ReplicationController:
		out, err = c.kc.CoreV1().ReplicationControllers(c.ns).Get(t.ObjectMeta.Name, options)
		// Deployment
	case *extensions.Deployment:
		out, err = c.kc.ExtensionsV1beta1().Deployments(c.ns).Get(t.ObjectMeta.Name, options)
	case *appsv1beta1.Deployment:
		out, err = c.kc.AppsV1beta1().Deployments(c.ns).Get(t.ObjectMeta.Name, options)
	case *appsv1beta2.Deployment:
		out, err = c.kc.AppsV1beta2().Deployments(c.ns).Get(t.ObjectMeta.Name, options)
	case *appsv1.Deployment:
		out, err = c.kc.AppsV1().Deployments(c.ns).Get(t.ObjectMeta.Name, options)
		// DaemonSet
	case *extensions.DaemonSet:
		out, err = c.kc.ExtensionsV1beta1().DaemonSets(c.ns).Get(t.ObjectMeta.Name, options)
	case *appsv1beta2.DaemonSet:
		out, err = c.kc.AppsV1beta2().DaemonSets(c.ns).Get(t.ObjectMeta.Name, options)
	case *appsv1.DaemonSet:
		out, err = c.kc.AppsV1().DaemonSets(c.ns).Get(t.ObjectMeta.Name, options)
		// ReplicaSet
	case *extensions.ReplicaSet:
		out, err = c.kc.ExtensionsV1beta1().ReplicaSets(c.ns).Get(t.ObjectMeta.Name, options)
	case *appsv1beta2.ReplicaSet:
		out, err = c.kc.AppsV1beta2().ReplicaSets(c.ns).Get(t.ObjectMeta.Name, options)
	case *appsv1.ReplicaSet:
		out, err = c.kc.AppsV1().ReplicaSets(c.ns).Get(t.ObjectMeta.Name, options)
		// StatefulSet
	case *appsv1beta1.StatefulSet:
		out, err = c.kc.AppsV1beta1().StatefulSets(c.ns).Get(t.ObjectMeta.Name, options)
	case *appsv1beta2.StatefulSet:
		out, err = c.kc.AppsV1beta2().StatefulSets(c.ns).Get(t.ObjectMeta.Name, options)
	case *appsv1.StatefulSet:
		out, err = c.kc.AppsV1().StatefulSets(c.ns).Get(t.ObjectMeta.Name, options)
		// Job
	case *batchv1.Job:
		out, err = c.kc.BatchV1().Jobs(c.ns).Get(t.ObjectMeta.Name, options)
		// CronJob
	case *batchv1beta1.CronJob:
		out, err = c.kc.BatchV1beta1().CronJobs(c.ns).Get(t.ObjectMeta.Name, options)
	case *ocapps.DeploymentConfig:
		out, err = c.oc.AppsV1().DeploymentConfigs(c.ns).Get(t.ObjectMeta.Name, options)
	default:
		err = fmt.Errorf("the object is not a pod or does not have a pod template")
	}
	if err != nil {
		return nil, err
	}
	return ConvertToWorkload(out)
}

func (c *workloads) List(opts metav1.ListOptions) (*v1.WorkloadList, error) {
	options := metav1.ListOptions{
		LabelSelector:        opts.LabelSelector,
		FieldSelector:        opts.FieldSelector,
		IncludeUninitialized: opts.IncludeUninitialized,
		ResourceVersion:      opts.ResourceVersion,
		TimeoutSeconds:       opts.TimeoutSeconds,
	}
	list := v1.WorkloadList{Items: make([]v1.Workload, 0)}

	if c.kc != nil {
		{
			objects, err := c.kc.AppsV1beta1().Deployments(c.ns).List(options)
			if err != nil {
				return nil, err
			}
			err = meta.EachListItem(objects, func(obj runtime.Object) error {
				w, err := ConvertToWorkload(obj)
				if err != nil {
					return err
				}
				list.Items = append(list.Items, *w)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
		{
			objects, err := c.kc.ExtensionsV1beta1().ReplicaSets(c.ns).List(options)
			if err != nil {
				return nil, err
			}
			err = meta.EachListItem(objects, func(obj runtime.Object) error {
				w, err := ConvertToWorkload(obj)
				if err != nil {
					return err
				}
				list.Items = append(list.Items, *w)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
		{
			if c.kc != nil {
				objects, err := c.kc.AppsV1beta1().StatefulSets(c.ns).List(options)
				if err != nil {
					return nil, err
				}
				err = meta.EachListItem(objects, func(obj runtime.Object) error {
					w, err := ConvertToWorkload(obj)
					if err != nil {
						return err
					}
					list.Items = append(list.Items, *w)
					return nil
				})
				if err != nil {
					return nil, err
				}
			}
		}
		{
			objects, err := c.kc.ExtensionsV1beta1().DaemonSets(c.ns).List(options)
			if err != nil {
				return nil, err
			}
			err = meta.EachListItem(objects, func(obj runtime.Object) error {
				w, err := ConvertToWorkload(obj)
				if err != nil {
					return err
				}
				list.Items = append(list.Items, *w)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
		{
			objects, err := c.kc.CoreV1().ReplicationControllers(c.ns).List(options)
			if err != nil {
				return nil, err
			}
			err = meta.EachListItem(objects, func(obj runtime.Object) error {
				w, err := ConvertToWorkload(obj)
				if err != nil {
					return err
				}
				list.Items = append(list.Items, *w)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
		{
			objects, err := c.kc.BatchV1().Jobs(c.ns).List(options)
			if err != nil {
				return nil, err
			}
			err = meta.EachListItem(objects, func(obj runtime.Object) error {
				w, err := ConvertToWorkload(obj)
				if err != nil {
					return err
				}
				list.Items = append(list.Items, *w)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
		{
			objects, err := c.kc.BatchV1beta1().CronJobs(c.ns).List(options)
			if err != nil {
				return nil, err
			}
			err = meta.EachListItem(objects, func(obj runtime.Object) error {
				w, err := ConvertToWorkload(obj)
				if err != nil {
					return err
				}
				list.Items = append(list.Items, *w)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
	}
	{
		if c.oc != nil {
			objects, err := c.oc.AppsV1().DeploymentConfigs(c.ns).List(options)
			if err != nil {
				return nil, err
			}
			err = meta.EachListItem(objects, func(obj runtime.Object) error {
				w, err := ConvertToWorkload(obj)
				if err != nil {
					return err
				}
				list.Items = append(list.Items, *w)
				return nil
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return &list, nil
}

func (c *workloads) Patch(cur *v1.Workload, transform WorkloadTransformerFunc) (*v1.Workload, kutil.VerbType, error) {
	mod := transform(cur.DeepCopy())
	err := ApplyWorkload(mod.Object, mod)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}
	return c.PatchObject(cur, mod)
}

func (c *workloads) PatchObject(cur, mod *v1.Workload) (*v1.Workload, kutil.VerbType, error) {
	curJson, err := json.Marshal(cur.Object)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}

	modJson, err := json.Marshal(mod.Object)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}

	patch, err := jsonpatch.CreateMergePatch(curJson, modJson)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}
	if len(patch) == 0 || string(patch) == "{}" {
		return cur, kutil.VerbUnchanged, nil
	}
	glog.V(3).Infof("Patching workload %s/%s with %s.", cur.Namespace, cur.Name, string(patch))

	var out runtime.Object
	switch mod.Object.(type) {
	case *core.Pod:
		out, err = c.kc.CoreV1().Pods(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
		// ReplicationController
	case *core.ReplicationController:
		out, err = c.kc.CoreV1().ReplicationControllers(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
		// Deployment
	case *extensions.Deployment:
		out, err = c.kc.ExtensionsV1beta1().Deployments(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *appsv1beta1.Deployment:
		out, err = c.kc.AppsV1beta1().Deployments(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *appsv1beta2.Deployment:
		out, err = c.kc.AppsV1beta2().Deployments(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *appsv1.Deployment:
		out, err = c.kc.AppsV1().Deployments(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
		// DaemonSet
	case *extensions.DaemonSet:
		out, err = c.kc.ExtensionsV1beta1().DaemonSets(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *appsv1beta2.DaemonSet:
		out, err = c.kc.AppsV1beta2().DaemonSets(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *appsv1.DaemonSet:
		out, err = c.kc.AppsV1().DaemonSets(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
		// ReplicaSet
	case *extensions.ReplicaSet:
		out, err = c.kc.ExtensionsV1beta1().ReplicaSets(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *appsv1beta2.ReplicaSet:
		out, err = c.kc.AppsV1beta2().ReplicaSets(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *appsv1.ReplicaSet:
		out, err = c.kc.AppsV1().ReplicaSets(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
		// StatefulSet
	case *appsv1beta1.StatefulSet:
		out, err = c.kc.AppsV1beta1().StatefulSets(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *appsv1beta2.StatefulSet:
		out, err = c.kc.AppsV1beta2().StatefulSets(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *appsv1.StatefulSet:
		out, err = c.kc.AppsV1().StatefulSets(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
		// Job
	case *batchv1.Job:
		out, err = c.kc.BatchV1().Jobs(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
		// CronJob
	case *batchv1beta1.CronJob:
		out, err = c.kc.BatchV1beta1().CronJobs(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	case *ocapps.DeploymentConfig:
		out, err = c.oc.AppsV1().DeploymentConfigs(c.ns).Patch(cur.Name, types.StrategicMergePatchType, patch)
	default:
		err = fmt.Errorf("the object is not a pod or does not have a pod template")
	}
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}
	result, err := ConvertToWorkload(out)
	return result, kutil.VerbPatched, err
}

func (c *workloads) CreateOrPatch(obj runtime.Object, transform WorkloadTransformerFunc) (*v1.Workload, kutil.VerbType, error) {
	gvk := obj.GetObjectKind().GroupVersionKind()
	if gvk.String() == "" {
		return nil, kutil.VerbUnchanged, fmt.Errorf("obj missing GroupVersionKind")
	}

	cur, err := c.Get(obj, metav1.GetOptions{})
	if kerr.IsNotFound(err) {
		name, err := meta.NewAccessor().Name(obj)
		if err != nil {
			return nil, kutil.VerbUnchanged, err
		}
		glog.V(3).Infof("Creating %s %s/%s.", gvk, c.ns, name)
		out, err := c.Create(transform(&v1.Workload{
			TypeMeta: metav1.TypeMeta{
				Kind:       gvk.Kind,
				APIVersion: gvk.GroupVersion().String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: c.ns,
				Name:      name,
			},
		}))
		return out, kutil.VerbCreated, err
	} else if err != nil {
		return nil, kutil.VerbUnchanged, err
	}
	return c.Patch(cur, transform)
}
