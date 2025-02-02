package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.Provider = &definednetProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &definednetProvider{
			version: version,
		}
	}
}

type definednetProvider struct {
	version string
}

func (p *definednetProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "definednet"
	resp.Version = p.version
}

func (p *definednetProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *definednetProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *definednetProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *definednetProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
