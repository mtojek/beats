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

package mgr_osd_disk

import (
	"fmt"

	"github.com/elastic/beats/metricbeat/helper"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/mb/parse"
)

const (
	defaultScheme      = "https"
	defaultPath        = "/request"
	defaultQueryParams = "wait=1"

	cephPrefix = "df"
)

var (
	hostParser = parse.URLHostParserBuilder{
		DefaultScheme: defaultScheme,
		DefaultPath:   defaultPath,
		QueryParams:   defaultQueryParams,
	}.Build()
)

func init() {
	mb.Registry.MustAddMetricSet("ceph", "mgr_osd_disk", New,
		mb.WithHostParser(hostParser),
	)
}

type MetricSet struct {
	mb.BaseMetricSet
	*helper.HTTP
}

func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	http, err := helper.NewHTTP(base)
	if err != nil {
		return nil, err
	}
	http.SetMethod("POST")
	http.SetHeader("Content-Type", "application/json")
	http.SetHeader("Accept", "application/json")
	http.SetBody([]byte(fmt.Sprintf(`{"prefix": "%s", "format": "json"}`, cephPrefix)))

	return &MetricSet{
		base,
		http,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {
	content, err := m.HTTP.FetchContent()
	if err != nil {
		return err
	}

	events, err := eventsMapping(content)
	if err != nil {
		return err
	}

	for _, event := range events {
		reported := reporter.Event(mb.Event{MetricSetFields: event})
		if !reported {
			m.Logger().Debug("error reporting event")
			return nil
		}
	}
	return nil
}
