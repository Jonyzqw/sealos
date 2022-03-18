/*
Copyright 2022 cuisongliu@qq.com.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	OpTest    OpType = "test"
	OpRemove  OpType = "remove"
	OpAdd     OpType = "add"
	OpReplace OpType = "replace"
	OpMove    OpType = "move"
	OpCopy    OpType = "copy"
)

type OpType string

type Patch struct {
	Op    OpType          `json:"op"`
	Path  string          `json:"path"`
	From  string          `json:"from,omitempty"`
	Value runtime.Unknown `json:"value,omitempty"`
}

// KubeadmSpec defines the desired state of Config
type KubeadmSpec struct {
	InitConfig      []Patch `json:"initConfig,omitempty"`
	ClusterConfig   []Patch `json:"clusterConfig,omitempty"`
	KubeProxyConfig []Patch `json:"kubeProxyConfig,omitempty"`
	KubeletConfig   []Patch `json:"kubeletConfig,omitempty"`
	JoinConfig      []Patch `json:"joinConfig,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kubeadm is the Schema for the configs API
type Kubeadm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KubeadmSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeadmList contains a list of Config
type KubeadmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kubeadm `json:"items"`
}