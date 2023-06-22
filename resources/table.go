package resources

import (
	"context"
	"fmt"

	"github.com/TheJumpCloud/jcapi"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func UsersTable() *schema.Table {
	return &schema.Table{
		Name:      "jumpcloud_users",
		Resolver:  fetchUsers,
		Transform: transformers.TransformWithStruct(&jcapi.JCUser{}),
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	return fmt.Errorf("not implemented")
}
