package provider

import (
	"context"
	"fmt"
	"terraform-provider-eas/internal/client"
	androidappcredentials "terraform-provider-eas/provider/android/appcredentials"
	"terraform-provider-eas/provider/android/googleserviceaccountkey"
	"terraform-provider-eas/provider/app"
	appleappcredentials "terraform-provider-eas/provider/apple/appcredentials"
	"terraform-provider-eas/provider/apple/appidentifier"
	"terraform-provider-eas/provider/apple/appstoreapikey"
	"terraform-provider-eas/provider/apple/certificate"
	"terraform-provider-eas/provider/apple/provisioningprofile"
	"terraform-provider-eas/provider/apple/pushkey"
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
			"eas_app_store_api_key":          appstoreapikey.DataSource(),
			"eas_apple_certificate":          certificate.DataSource(),
			"eas_apple_push_key":             pushkey.DataSource(),
			"eas_google_service_account_key": googleserviceaccountkey.DataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"eas_android_app_credentials": androidappcredentials.Resource(),
			"eas_app":                     app.Resource(),
			"eas_app_variable":            appvariable.Resource(),
			"eas_provisioning_profile":    provisioningprofile.Resource(),
			"eas_apple_app_identifier":    appidentifier.Resource(),
			"eas_apple_app_credentials":   appleappcredentials.Resource(),
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
