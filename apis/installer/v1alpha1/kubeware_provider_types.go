/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindKubewareProvider = "KubewareProvider"
	ResourceKubewareProvider     = "kubewarperovider"
	ResourceKubewareProviders    = "kubewarperoviders"
)

// KubewareProvider defines the schama for KubewareProvider operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
type KubewareProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubewareProviderSpec `json:"spec,omitempty"`
}

// KubewareProviderSpec is the schema for Operator Operator values file
type KubewareProviderSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string    `json:"fullnameOverride"`
	RegistryFQDN     string    `json:"registryFQDN"`
	ReplicaCount     int32     `json:"replicaCount"`
	Operator         Container `json:"operator"`
	RbacProxy        Container `json:"rbacproxy"`
	ImagePullPolicy  string    `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	PodLabels map[string]string `json:"podLabels"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	// +optional
	Apiserver  KubewareProviderApiserver `json:"apiserver"`
	Monitoring Monitoring                `json:"monitoring"`
	Provider   ProviderConfig            `json:"provider"`
}

type ProviderConfig struct {
	NamespacePrefix    string   `json:"namespacePrefix"`
	ProviderPrettyName string   `json:"providerPrettyName"`
	ConsumerScope      string   `json:"consumerScope"`
	External           External `json:"external"`
}

type External struct {
	Address    string `json:"address"`
	ServerName string `json:"serverName"`
	CAFile     string `json:"CAFile"`
}

type KubewareProviderApiserver struct {
	GroupPriorityMinimum        int             `json:"groupPriorityMinimum"`
	VersionPriority             int             `json:"versionPriority"`
	EnableMutatingWebhook       bool            `json:"enableMutatingWebhook"`
	EnableValidatingWebhook     bool            `json:"enableValidatingWebhook"`
	Ca                          string          `json:"ca"`
	BypassValidatingWebhookXray bool            `json:"bypassValidatingWebhookXray"`
	UseKubeapiserverFqdnForAks  bool            `json:"useKubeapiserverFqdnForAks"`
	Healthcheck                 HealthcheckSpec `json:"healthcheck"`
	ServingCerts                ServingCerts    `json:"servingCerts"`
}

type HealthcheckSpec struct {
	//+optional
	Enabled bool `json:"enabled"`
}

type ServingCerts struct {
	Generate bool `json:"generate"`
	//+optional
	CaCrt string `json:"caCrt"`
	//+optional
	ServerCrt string `json:"serverCrt"`
	//+optional
	ServerKey string `json:"serverKey"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubewareProviderList is a list of KubewareProviders
type KubewareProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubewareProvider CRD objects
	Items []KubewareProvider `json:"items,omitempty"`
}
