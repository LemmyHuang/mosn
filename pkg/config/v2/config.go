/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v2

import (
	"encoding/json"
	"github.com/c2h5oh/datasize"
)

// MOSNConfig make up mosn to start the mosn project
// Servers contains the listener, filter and so on
// ClusterManager used to manage the upstream
type MOSNConfig struct {
	Servers              []ServerConfig       `json:"servers,omitempty"`                //server config
	ClusterManager       ClusterManagerConfig `json:"cluster_manager,omitempty"`        //cluster config
	DisableUpgrade       bool                 `json:"disable_upgrade,omitempty"`        // not upgrade from old mosn, default false
	CloseGraceful        bool                 `json:"close_graceful,omitempty"`         // deprecated, same as disable_upgrade
	InheritOldMosnconfig bool                 `json:"inherit_old_mosnconfig,omitempty"` // inherit old mosn config switch, default false
	Tracing              TracingConfig        `json:"tracing,omitempty"`
	Metrics              MetricsConfig        `json:"metrics,omitempty"`
	RawDynamicResources  json.RawMessage      `json:"dynamic_resources,omitempty"` //dynamic_resources raw message
	RawStaticResources   json.RawMessage      `json:"static_resources,omitempty"`  //static_resources raw message
	RawAdmin             *Admin               `json:"admin,omitempty"`             // admin
	Debug                PProfConfig          `json:"pprof,omitempty"`
	Pid                  string               `json:"pid,omitempty"`                 // pid file
	UDSDir               string               `json:"uds_dir,omitempty"`             // unix domain socket directory
	Plugin               PluginConfig         `json:"plugin,omitempty"`              // plugin config
	ThirdPartCodec       ThirdPartCodecConfig `json:"third_part_codec,omitempty"`    // third part codec config
	Extends              []ExtendConfig       `json:"extends,omitempty"`             // extend config
	Wasms                []WasmPluginConfig   `json:"wasm_global_plugins,omitempty"` // wasm config
}

// PProfConfig is used to start a pprof server for debug
type PProfConfig struct {
	StartDebug bool `json:"debug"`      // If StartDebug is true, start a pprof, default is false
	Port       int  `json:"port_value"` // If port value is 0, will use 9090 as default
}

// Tracing configuration for a server
type TracingConfig struct {
	Enable bool                   `json:"enable,omitempty"`
	Tracer string                 `json:"tracer,omitempty"` // DEPRECATED
	Driver string                 `json:"driver,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
}

// MetricsConfig for metrics sinks
type MetricsConfig struct {
	SinkConfigs  []Filter          `json:"sinks"`
	StatsMatcher StatsMatcher      `json:"stats_matcher"`
	ShmZone      string            `json:"shm_zone"`
	ShmSize      datasize.ByteSize `json:"shm_size"`
	FlushMosn    bool              `json:"flush_mosn"`
	LazyFlush    bool              `json:"lazy_flush"`
}

// PluginConfig for plugin config
type PluginConfig struct {
	LogBase string `json:"log_base"`
}

// ThirdPartCodecType represents type of a third part codec
type ThirdPartCodecType string

// Third part codec consts
const (
	GoPlugin ThirdPartCodecType = "go-plugin"
	Wasm     ThirdPartCodecType = "wasm"
)

// ThirdPartCodec represents configuration for a third part codec
type ThirdPartCodec struct {
	Enable         bool                   `json:"enable,omitempty"`
	Type           ThirdPartCodecType     `json:"type,omitempty"`
	Path           string                 `json:"path,omitempty"`
	LoaderFuncName string                 `json:"loader_func_name,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
}

// ThirdPartCodecConfig represents configurations for third part codec
type ThirdPartCodecConfig struct {
	Codecs []ThirdPartCodec `json:"codecs"`
}

// ExtendConfig for any extends
type ExtendConfig struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

// StatsMatcher is a configuration for disabling stat instantiation.
// TODO: support inclusion_list
// TODO: support exclusion list/inclusion_list as pattern
type StatsMatcher struct {
	RejectAll       bool     `json:"reject_all,omitempty"`
	ExclusionLabels []string `json:"exclusion_labels,omitempty"`
	ExclusionKeys   []string `json:"exclusion_keys,omitempty"`
}

// Mode is mosn's starting type
type Mode uint8

// File means start from config file
// Xds means start from xds
// Mix means start both from file and Xds
const (
	File Mode = iota
	Xds
	Mix
)

func (c *MOSNConfig) Mode() Mode {
	if len(c.Servers) > 0 {
		if len(c.RawStaticResources) == 0 || len(c.RawDynamicResources) == 0 {
			return File
		}

		return Mix
	}
	if len(c.RawStaticResources) > 0 && len(c.RawDynamicResources) > 0 {
		return Xds
	}

	return File
}

type Admin struct {
	Address *AddressInfo `json:"address,omitempty"`
}

func (admin *Admin) GetAddress() string {
	if admin.Address == nil {
		return ""
	}
	return admin.Address.SocketAddress.Address
}

func (admin *Admin) GetPortValue() uint32 {
	if admin.Address == nil {
		return 0
	}
	return admin.Address.SocketAddress.PortValue
}

type AddressInfo struct {
	SocketAddress SocketAddress `json:"socket_address,omitempty"`
}

type SocketAddress struct {
	Address   string `json:"address,omitempty"`
	PortValue uint32 `json:"port_value,omitempty"`
}

func (c *MOSNConfig) GetAdmin() *Admin {
	return c.RawAdmin
}
