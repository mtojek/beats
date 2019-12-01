// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package activemq

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "activemq", asset.ModuleFieldsPri, AssetActivemq); err != nil {
		panic(err)
	}
}

// AssetActivemq returns asset data.
// This is the base64 encoded gzipped contents of module/activemq.
func AssetActivemq() string {
	return "eJyskL9Ow0AMxvc8xadOMDQPkAGJha1ISDAj6+Kkp15yqe0r6tujSymkNFCB6i3O+fvzW2LD+wrkzO+42xaAeQtcYXFcLQqgZnXiB/Oxr3BXAMAq1ikwmigYSNT3Le7Hi9UTQmzR+MBaFkDjOdRajUdL9NTxiV0e2w9coZWYho/NjOGp0lTN1sJUf66Pehvev0WZ7mdVD/M8asDWZGi5ZyHj/Mm5TJvb8Y57K8/MKdXezrynXS44P4yt0EjsvgiOqtlay8nj7wCmOZKynPz4GcOFQHlelOUfMEJsr49iGIJ3lN//CYijEK6J5JE6RmxGDAdteNWUaUzJCG8Tq+HGBVJFFAhrTOL4tpzNqUZu82pCjn8N+x4AAP//jyECmQ=="
}
