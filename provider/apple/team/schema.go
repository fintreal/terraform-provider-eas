package team

import (
	"terraform-provider-eas/provider/apple/team/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"name": {
				Description: "The name of the Apple Team",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"identifier": {
				Description: "The identifier of the Apple Team",
				Type:        schema.TypeString,
				Required:    true,
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
				Description:  "Apple Team type. Can be `COMPANY_OR_ORGANIZATION`, `IN_HOUSE` or `INDIVIDUAL`.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"COMPANY_OR_ORGANIZATION", "IN_HOUSE", "INDIVIDUAL"}, false),
			},
		},
	}
}
