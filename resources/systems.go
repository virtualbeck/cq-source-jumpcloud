package resources

import (
	"context"
	"fmt"

	"github.com/TheJumpCloud/jcapi"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/virtualbeck/cq-source-jumpcloud/client"
)

func SystemsTable() *schema.Table {
	return &schema.Table{
		Name:      "jumpcloud_systems",
		Resolver:  fetchSystems,
		Transform: transformers.TransformWithStruct(&jcapi.JCSystem{}, transformers.WithPrimaryKeys("Id")),
	}
}

func fetchSystems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	systemsList, err := c.JumpCloud.GetSystems(false)
	if err != nil {
		return fmt.Errorf("Could not read systems, err='%s'\n", err)
	}

	for _, system := range systemsList {
		res <- system
	}

	return nil
}
