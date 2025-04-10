package appstoreapikey

import (
	"terraform-provider-eas/provider/apple/appstoreapikey/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS Id of the App Store Api Key",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "EAS Name of the App Store Api Key",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"identifier": {
				Description: "Identifier of the App Store Api Key",
				Type:        schema.TypeString,
				Required:    true,
			},
			"issuer_identifier": {
				Description: "Issuer Identifier of the App Store Api Key",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}
