package provider

import (
	"context"
	"terraform-provider-eas/internal/client"
	"terraform-provider-eas/provider/app"
	"terraform-provider-eas/provider/appvariable"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Description: "expo personal access token or robot access token",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EXPO_TOKEN", ""),
			},
			"account_name": {
				Description: "expo user/organization account name",
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EXPO_ACCOUNT_NAME", ""),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"eas_app":          app.DataSource(),
			"eas_app_variable": appvariable.DataSource(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"eas_app":          app.Resource(),
			"eas_app_variable": appvariable.Resource(),
		},
		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
			token := d.Get("token").(string)
			accountName := d.Get("account_name").(string)

			client, err := client.NewEASClient(token, accountName)

			var diags diag.Diagnostics

			if err != nil {
				return nil, diags
			}

			return client, diags
		},
	}
}
