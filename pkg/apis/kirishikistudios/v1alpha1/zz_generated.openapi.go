// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/mackerel-operator/pkg/apis/kirishikistudios/v1alpha1.Mackerel":       schema_pkg_apis_kirishikistudios_v1alpha1_Mackerel(ref),
		"github.com/mackerel-operator/pkg/apis/kirishikistudios/v1alpha1.MackerelSpec":   schema_pkg_apis_kirishikistudios_v1alpha1_MackerelSpec(ref),
		"github.com/mackerel-operator/pkg/apis/kirishikistudios/v1alpha1.MackerelStatus": schema_pkg_apis_kirishikistudios_v1alpha1_MackerelStatus(ref),
	}
}

func schema_pkg_apis_kirishikistudios_v1alpha1_Mackerel(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Mackerel is the Schema for the mackerels API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mackerel-operator/pkg/apis/kirishikistudios/v1alpha1.MackerelSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mackerel-operator/pkg/apis/kirishikistudios/v1alpha1.MackerelStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/mackerel-operator/pkg/apis/kirishikistudios/v1alpha1.MackerelSpec", "github.com/mackerel-operator/pkg/apis/kirishikistudios/v1alpha1.MackerelStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kirishikistudios_v1alpha1_MackerelSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MackerelSpec defines the desired state of Mackerel",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_kirishikistudios_v1alpha1_MackerelStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MackerelStatus defines the observed state of Mackerel",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
