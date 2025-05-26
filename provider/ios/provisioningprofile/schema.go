package provisioningprofile

import (
	"terraform-provider-eas/provider/ios/provisioningprofile/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		DeleteContext: operations.Delete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS Provisioning Profile Id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"app_identifier_id": {
				Description: "EAS Apple App Identifier Id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"base64": {
				Description: "Base64 encoded Provisioning Profile file",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Sensitive:   true,
			},
		},
	}
}
