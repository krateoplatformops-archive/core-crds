package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Link to Krateo deployment plugins
type Value struct {
	// Name: property name
	Name string `json:"name"`

	// Value: property value
	Value string `json:"value"`
}

type DeploymentSpec struct {
	// Title:
	// +optional
	Title string `json:"title"`

	// Icon:
	// +optional
	Icon string `json:"icon,omitempty"`

	// Owner:
	// +optional
	Owner string `json:"owner,omitempty"`

	// ChartUrl
	ChartUrl string `json:"chartUrl"`

	// Values:
	// +optional
	Values []Value `json:"values,omitempty"`
}

// +kubebuilder:object:root=true

// A Deployment is a Krateo deployment API type.
// +kubebuilder:resource:scope=Cluster,categories={managed,krateo,deployments}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec DeploymentSpec `json:"spec"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployment
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentList `json:"items"`
}

// Deployment type metadata.
var (
	DeploymentKind             = reflect.TypeOf(Deployment{}).Name()
	DeploymentGroupKind        = schema.GroupKind{Group: Group, Kind: DeploymentKind}.String()
	DeploymentKindAPIVersion   = DeploymentKind + "." + SchemeGroupVersion.String()
	DeploymentGroupVersionKind = SchemeGroupVersion.WithKind(DeploymentKind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
