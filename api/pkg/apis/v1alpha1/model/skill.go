/*
Copyright 2022 The COA Authors
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

package model

// +kubebuilder:object:generate=true
type SkillSpec struct {
	DisplayName string            `json:"displayName,omitempty"`
	Parameters  map[string]string `json:"parameters,omitempty"`
	Nodes       []NodeSpec        `json:"nodes"`
	Properties  map[string]string `json:"properties,omitempty"`
	Bindings    []BindingSpec     `json:"bindings,omitempty"`
	Edges       []EdgeSpec        `json:"edges"`
}

// +kubebuilder:object:generate=true
type SkillPackageSpec struct {
	DisplayName string            `json:"displayName,omitempty"`
	Skill       string            `json:"skill"`
	Properties  map[string]string `json:"properties,omitempty"`
	Constraints string            `json:"constraints,omitempty"`
	Routes      []RouteSpec       `json:"routes,omitempty"`
}
