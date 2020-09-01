/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/suzerain-io/pinniped/generated/1.18/client/clientset/versioned/typed/idp/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeIDPV1alpha1 struct {
	*testing.Fake
}

func (c *FakeIDPV1alpha1) WebhookIdentityProviders(namespace string) v1alpha1.WebhookIdentityProviderInterface {
	return &FakeWebhookIdentityProviders{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeIDPV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
