package v1alpha1

import (
	"github.com/martin-helmich/kubernetes-crd-example/api/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type ProjectInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.ProjectList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Project, error)
	Create(*v1alpha1.Project) (*v1alpha1.Project, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	// ...
}

type projectClient struct {
	restClient rest.Interface
	ns         string
}

func (c *projectClient) List(opts metav1.ListOptions) (*v1alpha1.ProjectList, error) {
	result := v1alpha1.ProjectList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("projects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *projectClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.Project, error) {
	result := v1alpha1.Project{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("projects").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

func (c *projectClient) Create(project *v1alpha1.Project) (*v1alpha1.Project, error) {
	result := v1alpha1.Project{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("projects").
		Body(project).
		Do().
		Into(&result)

	return &result, err
}

func (c *projectClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("projects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}
