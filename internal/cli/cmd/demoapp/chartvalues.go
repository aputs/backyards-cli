// Copyright © 2019 Banzai Cloud
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

package demoapp

import (
	"github.com/banzaicloud/backyards-cli/pkg/helm"
)

type Values struct {
	Replicas       int        `json:"replicas"`
	Image          helm.Image `json:"image"`
	Services       bool       `json:"services"`
	IstioResources bool       `json:"istioresources"`
	Analytics      bool       `json:"analytics"`
	Bookings       bool       `json:"bookings"`
	Catalog        bool       `json:"catalog"`
	Frontpage      bool       `json:"frontpage"`
	MoviesV1       bool       `json:"moviesv1"`
	MoviesV2       bool       `json:"moviesv2"`
	MoviesV3       bool       `json:"moviesv3"`
	Notifications  bool       `json:"notifications"`
	Payments       bool       `json:"payments"`

	UseNamespaceResource bool `json:"useNamespaceResource"`
}
