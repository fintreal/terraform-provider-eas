// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/fintreal/expo-eas-sdk-go/eas"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &easProvider{}

type easProvider struct {
	client    eas.EASClient
	accountId string
}

type easProviderModel struct {
	Token       types.String `tfsdk:"token"`
	AccountName types.String `tfsdk:"account_name"`
}

func (p *easProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "eas"
}

func (p *easProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"token": schema.StringAttribute{
			Optional:  true,
			Sensitive: true,
		},
		"account_name": schema.StringAttribute{
			Optional: true,
		},
	}}
}

func (p *easProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config easProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Unknown 'token'",
			"The provider cannot create the Expo EAS client as there is an unknown configuration value for 'token'. "+
				"Either set the value statically in the configuration, or use the EXPO_TOKEN environment variable.",
		)
	}

	if config.AccountName.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("account_name"),
			"Unknown 'account'",
			"The provider cannot create the Expo EAS client as there is an unknown configuration value for 'account_name'. "+
				"Either set the value statically in the configuration, or use the EXPO_ACCOUNT_NAME environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	token := os.Getenv("EXPO_TOKEN")
	accountName := os.Getenv("EXPO_ACCOUNT_NAME")

	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	if !config.AccountName.IsNull() {
		accountName = config.AccountName.ValueString()
	}

	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Missing 'token'",
			"The provider cannot create the Expo EAS client as there is a missing or empty value for 'token'. "+
				"Set the 'token' value in the configuration or use the EXPO_TOKEN environment variable. ",
		)
	}

	if accountName == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("account_name"),
			"Missing 'account_name'",
			"The provider cannot create the Expo EAS API client as there is a missing or empty value for 'account_name'. "+
				"Set the 'account_name' value in the configuration or use the EXPO_TOKEN environment variable. ",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	easClient := eas.NewEASClient(token)
	account, err := easClient.Account.GetByName(accountName)

	if err != nil {
		resp.Diagnostics.AddError("Failed to create accountId for '"+accountName+"' account.", err.Error())
		return
	}

	client := &easProvider{
		client:    *easClient,
		accountId: account.Id,
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *easProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *easProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &easProvider{}
	}
}
