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

// defaultSubnetModel maps defaultSubnet schema data.
type defaultSubnetModel struct {
	VnetName    types.String `tfsdk:"vnet_name"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Range       types.String `tfsdk:"range"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &defaultSubnetDataSource{}
)

// NewVnetDataSource is a helper function to simplify the provider implementation.
func NewDefaultSubnetDataSource() datasource.DataSource {
	return &defaultSubnetDataSource{}
}

// defaultSubnetDataSource is the data source implementation.
type defaultSubnetDataSource struct{}

// Metadata returns the data source type name.
func (d *defaultSubnetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_default_subnet"
}

// Schema defines the schema for the data source.
func (d *defaultSubnetDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"vnet_name": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"range": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *defaultSubnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var _d defaultSubnetModel
	req.Config.Get(ctx, &_d)

	cdb := confdb.ConfDB{}
	cdb.Init(embed.FS{}, "", subscription, environment, region)
	defaultSubnet, _ := cdb.DefaultSubnet(_d.VnetName.ValueString())

	_json, _ := json.Marshal(defaultSubnet)
	_defaultSubnet := confdb.Subnet{}
	_ = json.Unmarshal(_json, &_defaultSubnet)

	// Map response body to model
	defaultSubnetState := defaultSubnetModel{
		VnetName:    types.StringValue(_d.VnetName.ValueString()),
		ID:          types.StringValue(_defaultSubnet.ID),
		Name:        types.StringValue(_defaultSubnet.Name),
		Description: types.StringValue("something something"),
		Range:       types.StringValue(_defaultSubnet.Range),
	}

	// Set state
	diags := resp.State.Set(ctx, &defaultSubnetState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
