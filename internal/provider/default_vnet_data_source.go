package provider

import (
	"context"
	"embed"
	"encoding/json"

	"codeberg.org/wrecking-yard/terraform-provider-confdb/internal/confdb"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
			"region": schema.StringAttribute{
				Computed: true,
			},
			"environment": schema.StringAttribute{
				Computed: true,
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
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *defaultVnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	cdb := confdb.ConfDB{}
	cdb.Init(embed.FS{}, "", subscription, environment, region)
	defaultVnet, _ := cdb.DefaultVNet()

	_json, _ := json.Marshal(defaultVnet)
	_defaultVnet := confdb.VNet{}
	_ = json.Unmarshal(_json, &_defaultVnet)

	// Map response body to model
	defaultVnetState := defaultVnetModel{
		ID:          types.StringValue(_defaultVnet.ID),
		Name:        types.StringValue(_defaultVnet.Name),
		Description: types.StringValue("something something"),
		Region:      types.StringValue(region),
		Environment: types.StringValue(environment),
	}

	// Set state
	diags := resp.State.Set(ctx, &defaultVnetState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
