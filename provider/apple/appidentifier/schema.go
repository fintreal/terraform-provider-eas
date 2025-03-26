package appidentifier

import (
	"terraform-provider-eas/provider/apple/appidentifier/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		DeleteContext: operations.Delete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS Apple App Identifier Id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"identifier": {
				Description: "Apple App Identifier",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}
