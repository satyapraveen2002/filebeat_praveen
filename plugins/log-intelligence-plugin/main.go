package main

import (
	_ "github.com/raboof/beats-output-http"

	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/plugin"
)

var Bundle = plugin.Bundle(
	outputs.Plugin("vmware-log-intelligence", vmware-log-intelligence.New),
)
