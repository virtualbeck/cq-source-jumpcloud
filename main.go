package main

import (
	"github.com/virtualbeck/cq-source-jumpcloud/plugin"

	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
