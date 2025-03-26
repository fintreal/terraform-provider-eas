package appcredentials

import (
	"strings"
	"terraform-provider-eas/provider/android/appcredentials/operations"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   operations.Read,
		CreateContext: operations.Create,
		DeleteContext: operations.Delete,
		UpdateContext: operations.Update,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "EAS Android App Credential Id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"app_id": {
				Description: "EAS App Id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"identifier": {
				Description: "Identifier of the Android App",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"google_service_account_key_id": {
				Description: "Google Service Account Key Id",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"fcm_key": {
				Description: "FCM Google Service Account Key Id",
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				ForceNew:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// Normalize whitespace by removing all whitespace and comparing
					oldNormalized := strings.Join(strings.Fields(old), "")
					newNormalized := strings.Join(strings.Fields(new), "")
					return oldNormalized == newNormalized
				},
			},
			"build_credentials": {
				Description: "EAS Android Build Credentials",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "EAS Android Build Credential Id",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"keystore_id": {
							Description: "EAS Id of the Android Keystore",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "Name of the Android Build Credential",
							Type:        schema.TypeString,
							Required:    true,
						},
					},
				},
			},
		},
	}
}
