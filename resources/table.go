package resources

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/TheJumpCloud/jcapi"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:     "cq-source-jumpcloud_sample_table",
		Resolver: fetchUsers,
		Transform: transformers.TransformWithStruct(&jcapi.JCUser{}),		
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	return fmt.Errorf("not implemented")
}
