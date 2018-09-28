// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}