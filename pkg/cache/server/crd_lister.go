/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"context"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionslisters "k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/kcp"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/clusters"

	"github.com/kcp-dev/kcp/pkg/cache/server/bootstrap"
)

// crdLister is a CRD lister
type crdLister struct {
	lister apiextensionslisters.CustomResourceDefinitionLister
}

var _ kcp.ClusterAwareCRDLister = &crdLister{}

// List lists all CustomResourceDefinitions
func (c *crdLister) List(ctx context.Context, selector labels.Selector) ([]*apiextensionsv1.CustomResourceDefinition, error) {
	// TODO: make it shard and cluster aware, for now just return what we have in the system ws
	return c.lister.List(selector)
}

func (c *crdLister) Refresh(crd *apiextensionsv1.CustomResourceDefinition) (*apiextensionsv1.CustomResourceDefinition, error) {
	return crd, nil
}

// Get gets a CustomResourceDefinition
func (c *crdLister) Get(ctx context.Context, name string) (*apiextensionsv1.CustomResourceDefinition, error) {
	// TODO: make it shard and cluster aware, for now just return what we have in the system ws
	return c.lister.Get(clusters.ToClusterAwareKey(bootstrap.SystemCRDLogicalCluster, name))
}
