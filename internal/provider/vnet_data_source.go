package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vnetDataSourceModel maps the data source schema data.
type vnetDataSourceModel struct {
	Vnet vnetModel `tfsdk:"vnet"`
}

// vnetModel maps vnet schema data.
type vnetModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &vnetDataSource{}
)

// NewVnetDataSource is a helper function to simplify the provider implementation.
func NewVnetDataSource() datasource.DataSource {
	return &vnetDataSource{}
}

// vnetDataSource is the data source implementation.
type vnetDataSource struct{}

// Metadata returns the data source type name.
func (d *vnetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vnet"
}

// Schema defines the schema for the data source.
func (d *vnetDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vnet": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *vnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Map response body to model
	vnetState := vnetDataSourceModel{
		Vnet: vnetModel{
			ID:          types.StringValue("/.../xyz/vnet-01"),
			Name:        types.StringValue("vnet-01"),
			Description: types.StringValue("some desc"),
		},
	}

	// Set state
	diags := resp.State.Set(ctx, &vnetState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
