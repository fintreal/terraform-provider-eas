package app

import (
	"fmt"
	"regexp"
	"terraform-provider-eas/provider/app/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "app id",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "display name",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"slug": {
				Description: "app slug",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func Resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: operations.Create,
		ReadContext:   operations.Read,
		UpdateContext: operations.Update,
		DeleteContext: operations.Delete,
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
