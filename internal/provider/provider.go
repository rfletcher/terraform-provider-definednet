package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-definednet/internal/definednet"
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

type definednetProviderModel struct {
	ApiKey types.String `tfsdk:"api_key"`
}

func (p *definednetProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "definednet"
	resp.Version = p.version
}

func (p *definednetProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *definednetProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config definednetProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Defined Networking API Key",
			"The provider cannot create the Defined Networking API client as there is an unknown configuration value for the definednet API key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the TF_DN_API_KEY environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// default to env variable; override with config
	// TODO: why does TF still prompt if the env var is set?
	api_key := os.Getenv("TF_DN_API_KEY")
	if !config.ApiKey.IsNull() {
		api_key = config.ApiKey.ValueString()
	}

	if api_key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing definednet API key",
			"The provider cannot create the Defined Networking API client as there is a missing or empty value for the API key. "+
				"Set the api_key value in the configuration or use the TF_DN_API_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	definednetClient, err := definednet.NewClient(api_key)

	if err != nil {
		return
	}

	resp.DataSourceData = definednetClient
	resp.ResourceData = definednetClient
}

func (p *definednetProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewHostsDataSource,
	}
}

func (p *definednetProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
