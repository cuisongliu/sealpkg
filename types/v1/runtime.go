// Copyright © 2023 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

type RuntimeConfigDefaultComponent struct {
	Containerd  string `json:"containerd"`
	Docker      string `json:"docker"`
	CRIDockerV3 string `json:"cri-docker-v3,omitempty"`
	CRIDockerV2 string `json:"cri-docker-v2,omitempty"`
	CRIDocker   string `json:"cri-docker,omitempty"`
	Sealos      string `json:"sealos"`
	CRIO        string `json:"crio"`
	Crun        string `json:"crun"`
	Runc        string `json:"runc"`
}

const (
	RuntimeK8s    string = "k8s"
	CRIDocker     string = "docker"
	CRIContainerd string = "containerd"
	CRICRIO       string = "crio"
)

type RuntimeConfigData struct {
	CRI            []string `json:"cri,omitempty"`
	Runtime        string   `json:"runtime"`
	RuntimeVersion []string `json:"runtimeVersion,omitempty"`
}

type RuntimeStatusConfigData struct {
	CRI            string `json:"cri,omitempty"`
	Runtime        string `json:"runtime"`
	RuntimeVersion string `json:"runtimeVersion,omitempty"`
}

type RuntimeStatus struct {
	*RuntimeConfigDefaultComponent `json:",inline"`
	*RuntimeStatusConfigData       `json:",inline"`
}

type RuntimeConfig struct {
	Config  *RuntimeConfigData             `json:"config,omitempty"`
	Default *RuntimeConfigDefaultComponent `json:"default,omitempty"`
}

type RuntimeStatusList struct {
	Include []RuntimeStatus `json:"include,omitempty"`
}
