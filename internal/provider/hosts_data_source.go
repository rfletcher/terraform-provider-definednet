package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"terraform-provider-definednet/internal/definednet"
)

var (
	_ datasource.DataSource              = &hostsDataSource{}
	_ datasource.DataSourceWithConfigure = &hostsDataSource{}
)

func NewHostsDataSource() datasource.DataSource {
	return &hostsDataSource{}
}

type hostsDataSource struct {
	Client *definednet.Client
}

type hostsDataSourceModel struct {
	Hosts []hostModel `tfsdk:"hosts"`
}

type hostModel struct {
	ID             types.String      `tfsdk:"id"`
	Name           types.String      `tfsdk:"name"`
	OrganizationId types.String      `tfsdk:"organization_id"`
	NetworkId      types.String      `tfsdk:"network_id"`
	RoleId         types.String      `tfsdk:"role_id"`
	IpAddress      types.String      `tfsdk:"ip_address"`
	ListenPort     types.Int64       `tfsdk:"listen_port"`
	IsBlocked      types.Bool        `tfsdk:"is_blocked"`
	IsLighthouse   types.Bool        `tfsdk:"is_lighthouse"`
	IsRelay        types.Bool        `tfsdk:"is_relay"`
	CreatedAt      types.String      `tfsdk:"created_at"`
	Metadata       hostMetadataModel `tfsdk:"metadata"`
	// StaticAddresses []types.String      `tfsdk:"static_addresses"` // TODO
}

type hostMetadataModel struct {
	LastSeenAt      types.String `tfsdk:"last_seen_at"`
	Version         types.String `tfsdk:"version"`
	Platform        types.String `tfsdk:"platform"`
	UpdateAvailable types.Bool   `tfsdk:"update_available"`
}

func (d *hostsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*definednet.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *definednet.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.Client = client
}

func (d *hostsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hosts"
}

func (d *hostsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"hosts": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"organization_id": schema.StringAttribute{
							Computed: true,
						},
						"network_id": schema.StringAttribute{
							Computed: true,
						},
						"role_id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"ip_address": schema.StringAttribute{
							Computed: true,
						},
						"listen_port": schema.Int64Attribute{
							Computed: true,
						},
						"is_lighthouse": schema.BoolAttribute{
							Computed: true,
						},
						"is_relay": schema.BoolAttribute{
							Computed: true,
						},
						"is_blocked": schema.BoolAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"metadata": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"last_seen_at": schema.StringAttribute{
									Computed: true,
								},
								"version": schema.StringAttribute{
									Computed: true,
								},
								"platform": schema.StringAttribute{
									Computed: true,
								},
								"update_available": schema.BoolAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *hostsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state hostsDataSourceModel

	hosts, err := d.Client.Hosts()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Defined Networking Hosts",
			err.Error(),
		)
		return
	}

	for _, host := range hosts {
		hostState := hostModel{
			ID:             types.StringValue(host.ID),
			Name:           types.StringValue(host.Name),
			OrganizationId: types.StringValue(host.OrganizationId),
			NetworkId:      types.StringValue(host.NetworkId),
			RoleId:         types.StringValue(host.RoleId),
			IpAddress:      types.StringValue(host.IpAddress),
			ListenPort:     types.Int64Value(int64(host.ListenPort)),
			IsBlocked:      types.BoolValue(host.IsBlocked),
			IsLighthouse:   types.BoolValue(host.IsLighthouse),
			IsRelay:        types.BoolValue(host.IsRelay),
			CreatedAt:      types.StringValue(host.CreatedAt),
			// StaticAddresses: []types.String      `tfsdk:"static_addresses"` // TODO
		}

		hostState.Metadata = hostMetadataModel{
			LastSeenAt:      types.StringValue(host.Metadata.LastSeenAt),
			Version:         types.StringValue(host.Metadata.Version),
			Platform:        types.StringValue(host.Metadata.Platform),
			UpdateAvailable: types.BoolValue(host.Metadata.UpdateAvailable),
		}

		state.Hosts = append(state.Hosts, hostState)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
