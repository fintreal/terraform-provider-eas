package googleserviceaccountkey

import (
	"terraform-provider-eas/provider/android/googleserviceaccountkey/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"project_identifier": {
				Description: "Project Identifier of the Google Service Account Key",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "EAS Id of the Google Service Account Key",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"client_email": {
				Description: "Google Service Account Email",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"client_identifier": {
				Description: "Google Service Account Identifier",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}
