package resources

import (
	"context"
	"fmt"
	"os"

	"github.com/TheJumpCloud/jcapi"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

var (
	apiURL      = config("JUMPCLOUD_API_URL", "https://console.jumpcloud.com/api")
	apiKey      = config("JUMPCLOUD_API_KEY", "")
	apiClientV1 = jcapi.NewJCAPI(apiKey, apiURL)
	isGroups    = true //IDK What this is
)

func UsersTable() *schema.Table {
	return &schema.Table{
		Name:      "jumpcloud_users",
		Resolver:  fetchUsers,
		Transform: transformers.TransformWithStruct(&jcapi.JCUser{}, transformers.WithPrimaryKeys("Id")),
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	userList, err := apiClientV1.GetSystemUsers(!isGroups)
	if err != nil {
		return fmt.Errorf("Could not read system users, err='%s'\n", err)
	}

	for _, user := range userList {
		res <- user
	}

	return nil
}

func config(s, e string) string {
	envVar := os.Getenv(s)
	if envVar != "" {
		return envVar
	}
	return e
}
