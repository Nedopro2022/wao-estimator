package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EstimatorSpec defines the desired state of Estimator
type EstimatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Estimator. Edit estimator_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// EstimatorStatus defines the observed state of Estimator
type EstimatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Estimator is the Schema for the estimators API
type Estimator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EstimatorSpec   `json:"spec,omitempty"`
	Status EstimatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EstimatorList contains a list of Estimator
type EstimatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Estimator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Estimator{}, &EstimatorList{})
}
