package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type SecretRef struct {
	// Name: secret name
	Name string `json:"name"`

	// Namespace: secret namespace
	Namespace string `json:"namespace,omitempty"`

	// Key: secret key
	Key string `json:"key,omitempty"`
}

type Repository struct {
	// Url: repository url
	Url string `json:"url"`

	// AuthMethod: authentication method
	AuthMethod string `json:"authMethod"`

	// SecretRef: secret reference
	// +optional
	SecretRef SecretRef `json:"secretRef,omitempty"`
}

// Link to Krateo gitops plugins
type Value struct {
	// Name: property name
	Name string `json:"name"`

	// Value: property value
	Value string `json:"value"`
}

type GitOpsSpec struct {
	// Title:
	// +optional
	Title string `json:"title"`

	// Icon:
	// +optional
	Icon string `json:"icon,omitempty"`

	// Owner:
	// +optional
	Owner string `json:"owner,omitempty"`

	// TargetRepo:
	TargetRepo Repository `json:"targetRepo"`

	// ChartUrl
	ChartUrl string `json:"chartUrl"`

	// Values:
	// +optional
	Values []Value `json:"values,omitempty"`
}

// +kubebuilder:object:root=true

// A GitOps is a Krateo gitops API type.
// +kubebuilder:resource:scope=Cluster,categories={managed,krateo,deployments}
type GitOps struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec GitOpsSpec `json:"spec"`
}

// +kubebuilder:object:root=true

// GitOpsList contains a list of GitOps
type GitOpsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitOpsList `json:"items"`
}

// GitOps type metadata.
var (
	GitOpsKind             = reflect.TypeOf(GitOps{}).Name()
	GitOpsGroupKind        = schema.GroupKind{Group: Group, Kind: GitOpsKind}.String()
	GitOpsKindAPIVersion   = GitOpsKind + "." + SchemeGroupVersion.String()
	GitOpsGroupVersionKind = SchemeGroupVersion.WithKind(GitOpsKind)
)

func init() {
	SchemeBuilder.Register(&GitOps{}, &GitOpsList{})
}
