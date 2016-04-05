// Copyright 2016 fatedier, fatedier@gmail.com
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

package msg

type GeneralRes struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

// messages between control connection of frpc and frps
type ControlReq struct {
	Type          int64  `json:"type"`
	ProxyName     string `json:"proxy_name,omitempty"`
	AuthKey       string `json:"auth_key, omitempty"`
	UseEncryption bool   `json:"use_encryption, omitempty"`
	Timestamp     int64  `json:"timestamp, omitempty"`
}

type ControlRes struct {
	Type int64  `json:"type"`
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
