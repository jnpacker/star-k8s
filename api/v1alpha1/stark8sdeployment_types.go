/*


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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StarK8sDeploymentSpec defines the desired state of StarK8sDeployment
type StarK8sDeploymentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	PowerState        string             `json:"powerState,omitempty"`
	Provider          Provider           `json:"provider"`
	ProviderSecretRef v1.SecretEnvSource `json:"providerSecretRef"`
}

type Provider struct {
	Eks *Eks `json:"EKS,omitempty"`
	//Aks *Aks `json:"AKS"`
	//Gke *Gke `json:"GKE"`
}

type Eks struct {
	ControlPlane ControlPlane `json:"controlPlane"`
	DataPlane    []DataPlane  `json:"dataPlane,omitempty"`
}

type ControlPlane struct {
	Nodes int                   `json:"nodes"`
	Spec  *runtime.RawExtension `json:"spec,omitempty"`
}

type DataPlane struct {
	Nodes int                   `json:"nodes"`
	Spec  *runtime.RawExtension `json:"spec,omitempty"`
}

// StarK8sDeploymentStatus defines the observed state of StarK8sDeployment
type StarK8sDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []v1.ComponentCondition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// StarK8sDeployment is the Schema for the stark8sdeployments API
type StarK8sDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StarK8sDeploymentSpec   `json:"spec,omitempty"`
	Status StarK8sDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StarK8sDeploymentList contains a list of StarK8sDeployment
type StarK8sDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StarK8sDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StarK8sDeployment{}, &StarK8sDeploymentList{})
}
