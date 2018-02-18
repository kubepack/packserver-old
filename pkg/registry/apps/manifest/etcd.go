package manifest

import (
	"github.com/kubepack/packserver/apis/apps"
	"github.com/kubepack/packserver/pkg/util/restoptions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

// REST contains the REST storage for Manifest objects.
type REST struct {
	*registry.Store
}

var _ rest.StandardStorage = &REST{}
var _ rest.ShortNamesProvider = &REST{}
var _ rest.CategoriesProvider = &REST{}

// Categories implements the CategoriesProvider interface. Returns a list of categories a resource is part of.
func (r *REST) Categories() []string {
	return []string{"all"}
}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"dc"}
}

// NewREST returns a deploymentConfigREST containing the REST storage for Manifest objects,
// a statusREST containing the REST storage for changing the status of a Manifest,
// and a scaleREST containing the REST storage for the Scale subresources of Packs.
func NewREST(optsGetter restoptions.Getter) (*REST, *StatusREST, error) {
	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &apps.Manifest{} },
		NewListFunc:              func() runtime.Object { return &apps.ManifestList{} },
		DefaultQualifiedResource: apps.Resource("deploymentconfigs"),

		CreateStrategy: GroupStrategy,
		UpdateStrategy: GroupStrategy,
		DeleteStrategy: GroupStrategy,
	}

	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, nil, err
	}

	deploymentConfigREST := &REST{store}

	statusStore := *store
	statusStore.UpdateStrategy = StatusStrategy
	statusREST := &StatusREST{store: &statusStore}

	return deploymentConfigREST, statusREST, nil
}

// StatusREST implements the REST endpoint for changing the status of a Manifest.
type StatusREST struct {
	store *registry.Store
}

// StatusREST implements Patcher
var _ = rest.Patcher(&StatusREST{})

func (r *StatusREST) New() runtime.Object {
	return &apps.Manifest{}
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *StatusREST) Get(ctx apirequest.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	return r.store.Get(ctx, name, options)
}

// Update alters the status subset of an deploymentConfig.
func (r *StatusREST) Update(ctx apirequest.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc) (runtime.Object, bool, error) {
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation)
}

// LegacyREST allows us to wrap and alter some behavior
type LegacyREST struct {
	*REST
}

func (r *LegacyREST) Categories() []string {
	return []string{}
}
