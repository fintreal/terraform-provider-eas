package app

import (
	"fmt"
	"regexp"
	"terraform-provider-eas/provider/app/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: operations.Create,
		ReadContext:   operations.Read,
		UpdateContext: operations.Update,
		DeleteContext: operations.Delete,
		Importer:      &schema.ResourceImporter{StateContext: operations.Import},
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "app id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "display name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"slug": {
				Description:  "app slug",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateSlug,
			},
		},
	}
}

func validateSlug(i any, s string) ([]string, []error) {
	matched, err := regexp.MatchString(`^[a-z0-9]+(-[a-z0-9]+)*$`, i.(string))
	if err != nil {
		return nil, []error{fmt.Errorf("error validating slug: %v", err)}
	}
	if !matched {
		return nil, []error{fmt.Errorf("slug must be lowercase alphanumeric and may contain single dashes between segments (no double dashes)")}
	}
	return nil, nil
}
