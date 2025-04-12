package pushkey

import (
	"terraform-provider-eas/provider/apple/pushkey/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS Id of the Apple Push Key",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"identifier": {
				Description: "Identifier of the Push Key",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
