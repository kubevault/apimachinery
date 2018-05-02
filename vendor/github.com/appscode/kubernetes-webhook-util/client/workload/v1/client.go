package v1

import (
	wpi "github.com/appscode/kubernetes-webhook-util/apis/workload/v1"
	"github.com/appscode/kutil/discovery"
	ocapps "github.com/openshift/api/apps/v1"
	occ "github.com/openshift/client-go/apps/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Interface interface {
	WorkloadsGetter
}

// Client is used to interact with features provided by the storage.k8s.io group.
type Client struct {
	kc kubernetes.Interface
	oc occ.Interface
}

func (c *Client) Workloads(namespace string) WorkloadInterface {
	return newWorkloads(c.kc, c.oc, namespace)
}

// NewForConfig creates a new Client for the given config.
func NewForConfig(c *rest.Config) (*Client, error) {
	kc, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	var oc occ.Interface
	if discovery.IsPreferredAPIResource(kc.Discovery(), ocapps.SchemeGroupVersion.String(), wpi.KindDeploymentConfig) {
		oc, err = occ.NewForConfig(c)
		if err != nil {
			return nil, err
		}
	}
	return &Client{kc, oc}, nil
}

// NewForConfigOrDie creates a new Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Client {
	kc := kubernetes.NewForConfigOrDie(c)
	var oc occ.Interface
	var err error
	if discovery.IsPreferredAPIResource(kc.Discovery(), ocapps.SchemeGroupVersion.String(), wpi.KindDeploymentConfig) {
		oc, err = occ.NewForConfig(c)
		if err != nil {
			panic(err)
		}
	}
	return &Client{kc, oc}
}

// New creates a new Client for the given RESTClient.
func New(kc kubernetes.Interface, oc occ.Interface) *Client {
	return &Client{kc, oc}
}
