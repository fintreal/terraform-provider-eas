package accountvariable

import (
	"strings"
	"terraform-provider-eas/provider/accountvariable/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Resource defines the eas_account_variable resource. Account environment
// variables belong to the account configured on the provider (account_name),
// so — unlike eas_app_variable — this resource takes no app/account id.
func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		UpdateContext: operations.Update,
		DeleteContext: operations.Delete,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The name of the account environment variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "The id of the environment variable",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"value": {
				Description: "The value of the account environment variable",
				Type:        schema.TypeString,
				Required:    true,
			},
			"visibility": {
				Description:  "The visibility of the account environment variable",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PUBLIC", "SENSITIVE", "SECRET"}, false),
			},
			"environments": {
				Description: "The environments of the account environment variable",
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
