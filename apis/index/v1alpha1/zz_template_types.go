/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TemplateObservation struct {

	// The name of the index template.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TemplateParameters struct {

	// The JSON body of the index template.
	// +kubebuilder:validation:Required
	Body *string `json:"body" tf:"body,omitempty"`

	// The name of the index template.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// TemplateSpec defines the desired state of Template
type TemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TemplateParameters `json:"forProvider"`
}

// TemplateStatus defines the observed state of Template.
type TemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Template is the Schema for the Templates API. Provides an Elasticsearch index template resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opensearch}
type Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TemplateSpec   `json:"spec"`
	Status            TemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TemplateList contains a list of Templates
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Template `json:"items"`
}

// Repository type metadata.
var (
	Template_Kind             = "Template"
	Template_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Template_Kind}.String()
	Template_KindAPIVersion   = Template_Kind + "." + CRDGroupVersion.String()
	Template_GroupVersionKind = CRDGroupVersion.WithKind(Template_Kind)
)

func init() {
	SchemeBuilder.Register(&Template{}, &TemplateList{})
}
