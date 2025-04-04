package appvariable

import (
	"terraform-provider-eas/provider/appvariable/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"app_id": {
				Description: "app id",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "environment variable name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "environment variable id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"value": {
				Description: "environment variable value",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"visibility": {
				Description: "visibility: PUBLIC, SENSITIVE, SECRET",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"environments": {
				Description: "environments: DEVELOPMENT, PREVIEW, PRODUCTION",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
					ValidateFunc: validation.StringInSlice([]string{"DEVELOPMENT", "PREVIEW", "PRODUCTION"}, false),
				},
				Required: true,
			},
		},
	}
}
