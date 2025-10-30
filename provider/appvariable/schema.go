package appvariable

import (
	"strings"
	"terraform-provider-eas/provider/appvariable/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		UpdateContext: operations.Update,
		DeleteContext: operations.Delete,
		Schema: map[string]*schema.Schema{
			"app_id": {
				Description: "The id of the app for the environment variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The name of the app for the environment variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "The id of the environment variable",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"value": {
				Description: "The value of the app for the environment variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"visibility": {
				Description:  "The visibility of the app for the environment variable",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PUBLIC", "SENSITIVE", "SECRET"}, false),
			},
			"environments": {
				Description: "The environments of the app for the environment variable",
				Type:        schema.TypeSet,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"development", "preview", "production"}, true),
					StateFunc: func(val any) string {
						return strings.ToLower(val.(string))
					},
				},
				Required: true,
			},
		},
	}
}
