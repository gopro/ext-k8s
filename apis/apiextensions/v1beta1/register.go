package v1beta1

import "github.com/gopro/ext-k8s"

func init() {
	k8s.Register("apiextensions.k8s.io", "v1beta1", "customresourcedefinitions", false, &CustomResourceDefinition{})

	k8s.RegisterList("apiextensions.k8s.io", "v1beta1", "customresourcedefinitions", false, &CustomResourceDefinitionList{})
}
