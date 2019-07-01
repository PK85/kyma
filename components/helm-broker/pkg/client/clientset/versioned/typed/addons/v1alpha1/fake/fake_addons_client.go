// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kyma-project/kyma/components/helm-broker/pkg/client/clientset/versioned/typed/addons/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAddonsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAddonsV1alpha1) AddonsConfigurations(namespace string) v1alpha1.AddonsConfigurationInterface {
	return &FakeAddonsConfigurations{c, namespace}
}

func (c *FakeAddonsV1alpha1) ClusterAddonsConfigurations() v1alpha1.ClusterAddonsConfigurationInterface {
	return &FakeClusterAddonsConfigurations{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAddonsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
