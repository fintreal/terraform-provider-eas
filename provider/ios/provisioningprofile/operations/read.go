package operations

import (
	"context"
	"terraform-provider-eas/internal/client"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*client.EASClient)

	input := eas.GetProvisioningProfileData{
		AccountId: client.AccountId,
		Id:        d.Get("id").(string),
	}
	data, err := client.Apple.ProvisioningProfile.Get(input)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)

	if err := d.Set("id", data.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("base64", data.Base64); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("app_identifier_id", data.AppBundleIdentifierId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
