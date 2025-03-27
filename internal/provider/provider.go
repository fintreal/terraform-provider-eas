// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &EASClient{}
var _ provider.ProviderWithFunctions = &EASClient{}
var _ provider.ProviderWithEphemeralResources = &EASClient{}

// EASClient defines the provider implementation.
type EASClient struct {
    token string
}

// EasProviderModel describes the provider data model.
type EasProviderModel struct {
	token types.String `tfsdk:"token"`
}

func (p *EASClient) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "eas"
}

func (p *EASClient) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"token": schema.StringAttribute{
				MarkdownDescription: "EXPO_TOKEN",
				Optional:            true,
			},
		},
	}
}

func (p *EASClient) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data EasProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	token := os.Getenv("EXPO_TOKEN")

	if !data.token.IsNull() {
		token = data.token.ValueString()
	}

	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Provider couldn't be created.",
			"Either provider 'token' or EXPO_TOKEN environment variable is required",
		)
	}

	client := &EASClient{token: token}
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *EASClient) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProjectResource,
	}
}

func (p *EASClient) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		NewExampleEphemeralResource,
	}
}

func (p *EASClient) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProjectDataSource,
	}
}

func (p *EASClient) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewExampleFunction,
	}
}

func New(token string) func() provider.Provider {
	return func() provider.Provider {
		return &EASClient{token: token}
	}
}
