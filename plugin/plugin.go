package plugin

import (
	"github.com/virtualbeck/cq-source-jumpcloud/client"
	"github.com/virtualbeck/cq-source-jumpcloud/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "v0.0.2"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"cq-source-jumpcloud",
		Version,
		schema.Tables{
			resources.UsersTable(),
			resources.SystemsTable(),
		},
		client.New,
	)
}
