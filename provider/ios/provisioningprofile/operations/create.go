package operations

import (
	"context"
	"terraform-provider-eas/internal/client"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)

	appIdentifierId := d.Get("app_identifier_id").(string)
	base64 := d.Get("base64").(string)

	input := eas.CreateProvisioningProfileData{
		AccountId:             client.AccountId,
		AppBundleIdentifierId: appIdentifierId,
		Base64:                base64,
	}

	data, err := client.Apple.ProvisioningProfile.Create(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)
	d.Set("id", data.Id)
	d.Set("app_identifier_id", data.AppBundleIdentifierId)
	d.Set("base64", data.Base64)

	var diags diag.Diagnostics
	return diags
}
