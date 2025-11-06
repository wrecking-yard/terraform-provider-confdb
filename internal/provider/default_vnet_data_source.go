package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// defaultVnetDataSourceModel maps the data source schema data.
type defaultVnetDataSourceModel struct {
	Vnet defaultVnetModel `tfsdk:"default_vnet"`
}

// defaultVnetModel maps defaultVnet schema data.
type defaultVnetModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Region      types.String `tfsdk:"region"`
	Environment types.String `tfsdk:"environment"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &defaultVnetDataSource{}
)

// NewVnetDataSource is a helper function to simplify the provider implementation.
func NewDefaultVnetDataSource() datasource.DataSource {
	return &defaultVnetDataSource{}
}

// defaultVnetDataSource is the data source implementation.
type defaultVnetDataSource struct{}

// Metadata returns the data source type name.
func (d *defaultVnetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_default_vnet"
}

// Schema defines the schema for the data source.
func (d *defaultVnetDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"default_vnet": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"region": schema.StringAttribute{
						Required: true,
					},
					"environment": schema.StringAttribute{
						Required: true,
					},
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
func (d *defaultVnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	fmt.Printf("xxx region: %s\n", region)
	fmt.Printf("xxx environment: %s\n", environment)
	fmt.Printf("xxx subscription: %s\n", subscription)

	// Map response body to model
	defaultVnetState := defaultVnetDataSourceModel{
		Vnet: defaultVnetModel{
			ID:          types.StringValue("/.../xyz/defaultVnet-01"),
			Name:        types.StringValue("defaultVnet-01"),
			Description: types.StringValue("some desc"),
			Region:      types.StringValue("some desc"),
			Environment: types.StringValue("some desc"),
		},
	}

	// Set state
	diags := resp.State.Set(ctx, &defaultVnetState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
