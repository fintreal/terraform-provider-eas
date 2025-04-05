package appleteam

import (
	"terraform-provider-eas/provider/appleteam/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the Apple Team",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_id": {
				Description: "The account ID associated with the Apple Team",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the Apple Team",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"identifier": {
				Description: "The identifier of the Apple Team",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description: "The account ID associated with the Apple Team",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		UpdateContext: operations.Update,
		DeleteContext: operations.Delete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The ID of the Apple Team",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"account_id": {
				Description: "The account ID associated with the Apple Team",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of the Apple Team",
				Type:        schema.TypeString,
				Required:    true,
			},
			"identifier": {
				Description: "The identifier of the Apple Team",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description: "The account ID associated with the Apple Team",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}
