// Copyright (C) 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import bgp "github.com/osrg/gobgp/config"

type VirtualNetwork struct {
	VNI                 uint32
	VxlanPort           uint16
	VtepInterface       string
	Color               uint32
	MemberInterfaces    []string
	ConfigureInterfaces []string
}

type Dataplane struct {
	Type               string
	VirtualNetworkList []VirtualNetwork
}

type Config struct {
	Bgp       bgp.Bgp
	Dataplane Dataplane
}
