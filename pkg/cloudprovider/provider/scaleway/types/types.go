/*
Copyright 2019 The Machine Controller Authors.

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

package types

import (
	providerconfigtypes "github.com/kubermatic/machine-controller/pkg/providerconfig/types"
)

type RawConfig struct {
	AccessKey      providerconfigtypes.ConfigVarString   `json:"accessKey,omitempty"`
	SecretKey      providerconfigtypes.ConfigVarString   `json:"secretKey,omitempty"`
	ProjectID      providerconfigtypes.ConfigVarString   `json:"projectId,omitempty"`
	Zone           providerconfigtypes.ConfigVarString   `json:"zone,omitempty"`
	CommercialType providerconfigtypes.ConfigVarString   `json:"commercialType"`
	IPv6           providerconfigtypes.ConfigVarBool     `json:"ipv6"`
	Tags           []providerconfigtypes.ConfigVarString `json:"tags,omitempty"`
}
