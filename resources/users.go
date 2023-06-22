package resources

import (
	"context"
	"fmt"

	"github.com/TheJumpCloud/jcapi"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/virtualbeck/cq-source-jumpcloud/client"
)

func UsersTable() *schema.Table {
	return &schema.Table{
		Name:      "jumpcloud_users",
		Resolver:  fetchUsers,
		Transform: transformers.TransformWithStruct(&jcapi.JCUser{}, transformers.WithPrimaryKeys("Id")),
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	userList, err := c.JumpCloud.GetSystemUsers(false)
	if err != nil {
		return fmt.Errorf("Could not read system users, err='%s'\n", err)
	}

	for _, user := range userList {
		res <- user
	}

	return nil
}
