package main

import (
  "os"
  "fmt"

  _ "github.com/raboof/beats-output-http"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/raboof/connbeat/beater"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/plugin"
  "vmware-log-intelligence"
)

var Bundle = plugin.Bundle(
	outputs.Plugin("VMware Log intelligence", vmware-log-intelligence.New),
)
