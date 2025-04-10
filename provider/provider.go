package provider

import (
	"context"
	"fmt"
	"terraform-provider-eas/internal/client"
	"terraform-provider-eas/provider/app"
	"terraform-provider-eas/provider/apple/appstoreapikey"
	"terraform-provider-eas/provider/apple/certificate"
	"terraform-provider-eas/provider/apple/provisioningprofile"
	"terraform-provider-eas/provider/apple/team"
	"terraform-provider-eas/provider/appvariable"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Description: "Expo personal access token or robot access token. You can set this via `EXPO TOKEN` environment variable.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EXPO_TOKEN", ""),
			},
			"account_name": {
				Description: "Expo user/organization account name. You can set this via `EXPO_ACCOUNT_NAME` environment variable.",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EXPO_ACCOUNT_NAME", ""),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"eas_app":                  app.DataSource(),
			"eas_app_variable":         appvariable.DataSource(),
			"eas_apple_team":           team.DataSource(),
			"eas_app_store_api_key":    appstoreapikey.DataSource(),
			"eas_apple_certificate":    certificate.DataSource(),
			"eas_provisioning_profile": provisioningprofile.DataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"eas_app":                  app.Resource(),
			"eas_app_variable":         appvariable.Resource(),
			"eas_apple_team":           team.Resource(),
			"eas_provisioning_profile": provisioningprofile.Resource(),
		},
		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
			token := d.Get("token").(string)
			accountName := d.Get("account_name").(string)

			var diags diag.Diagnostics

			if token == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "token value cannot be an empty string",
					Detail:   "set the token value in the provider configuration or via the EXPO_TOKEN environment variable",
				})
			}

			if accountName == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "account_name value cannot be an empty string",
					Detail:   "set the token value in the provider configuration or via the EXPO_ACCOUNT_NAME environment variable",
				})
			}

			if len(diags) > 0 {
				return nil, diags
			}

			client, err := client.NewEASClient(token, accountName)

			if err != nil {
				return nil, diag.Diagnostics{
					diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Failed to initialize EAS client",
						Detail:   fmt.Sprintf("Error: %v", err),
					},
				}
			}

			return client, diags
		},
	}
}
