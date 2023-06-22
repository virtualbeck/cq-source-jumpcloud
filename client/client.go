package client

import (
	"context"
	"fmt"
	"os"

	"github.com/TheJumpCloud/jcapi"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger    zerolog.Logger
	JumpCloud *jcapi.JCAPI
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "ID"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
	var (
		apiURL      = config("JUMPCLOUD_API_URL", "https://console.jumpcloud.com/api")
		apiKey      = config("JUMPCLOUD_API_KEY", "")
		apiClientV1 = jcapi.NewJCAPI(apiKey, apiURL)
	)

	return &Client{
		Logger:    logger,
		JumpCloud: &apiClientV1,
	}, nil
}
func config(s, e string) string {
	envVar := os.Getenv(s)
	if envVar != "" {
		return envVar
	}
	return e
}
