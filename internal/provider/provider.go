package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	// Ensure the implementation satisfies the expected interfaces.
	_            provider.Provider = &confDBProvider{}
	environment  string
	subscription string
	region       string
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &confDBProvider{
			version: version,
		}
	}
}

// confDBProvider is the provider implementation.
type confDBProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *confDBProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "confdb"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *confDBProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"subscription": schema.StringAttribute{
				Required: true,
			},
			"environment": schema.StringAttribute{
				Required: true,
			},
			"region": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

// ConfDBProviderModel maps provider schema data to a Go type.
type ConfDBProviderModel struct {
	Subscription types.String `tfsdk:"subscription"`
	Region       types.String `tfsdk:"region"`
	Environment  types.String `tfsdk:"environment"`
}

// Configure does nothing, at this point of time API client is not needed.
func (p *confDBProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// this could perhaps be passed as data source client parameters.
	// those attributes should not change within a single provider instance / terraform module
	// so i guess it's fine the way it is.
	var config ConfDBProviderModel
	_ = req.Config.Get(ctx, &config)
	environment = config.Environment.ValueString()
	subscription = config.Subscription.ValueString()
	region = config.Region.ValueString()
}

// Resources does nothing - at this point of time provider is not implementing any resources.
func (p *confDBProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

// DataSources defines the data sources implemented in the provider.
func (p *confDBProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewVnetDataSource,
		NewDefaultVnetDataSource,
	}
}
