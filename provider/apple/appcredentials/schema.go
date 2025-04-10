package appcredentials

import (
	"terraform-provider-eas/provider/apple/appcredentials/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: operations.Read,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS App Credentials Id",
				Type:        schema.TypeString,
				Required:    true,
			},
			"app_id": {
				Description: "EAS App Id",
				Type:        schema.TypeString,
				Required:    true,
			},
			"app_identifier_id": {
				Description: "EAS App Identifier Id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"app_store": {
				Description: "EAS App Build Credentials for App Store",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "EAS Build Credential Id",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"certificate_id": {
							Description: "EAS Id of the Apple Distribution Certificate",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"provisioning_profile_id": {
							Description: "EAS Provisioning Profile Id",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		DeleteContext: operations.Delete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS App Credentials Id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"app_id": {
				Description: "EAS App Id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"app_identifier_id": {
				Description: "EAS App Identifier Id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"app_store": {
				Description: "EAS App Build Credentials for App Store",
				Type:        schema.TypeList,
				ForceNew:    true,
				Optional:    true,
				MaxItems:    1, // ensures it's treated like an object
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "EAS Build Credential Id",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"certificate_id": {
							Description: "EAS Id of the Apple Distribution Certificate",
							Type:        schema.TypeString,
							Required:    true,
						},
						"provisioning_profile_id": {
							Description: "EAS Provisioning Profile Id",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}
}
