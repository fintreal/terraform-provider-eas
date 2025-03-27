// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &EASClient{}
var _ provider.ProviderWithFunctions = &EASClient{}
var _ provider.ProviderWithEphemeralResources = &EASClient{}

// EASClient defines the provider implementation.
type EASClient struct { }

// EasProviderModel describes the provider data model.
type EasProviderModel struct { }

func (p *EASClient) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "eas"
}

func (p *EASClient) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{}}
}

func (p *EASClient) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	client := &EASClient{}
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
		NewProjectVariableDataSource,
	}
}

func (p *EASClient) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewExampleFunction,
	}
}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &EASClient{}
	}
}
