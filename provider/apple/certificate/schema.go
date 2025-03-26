package certificate

import (
	"terraform-provider-eas/provider/apple/certificate/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS Id of the Apple Certificate",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"serial_number": {
				Description: "Apple Certificate serial number",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
