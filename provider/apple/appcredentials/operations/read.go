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

	input := eas.GetAppCredentialsData{
		Id:    d.Get("id").(string),
		AppId: d.Get("app_id").(string),
	}
	data, err := client.Apple.AppCredentials.Get(input)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)

	if err := d.Set("id", data.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("app_id", data.AppId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("app_identifier_id", data.AppIdentifierId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	var appStoreMap map[string]any
	for _, buildCredential := range data.BuildCredentials {
		if buildCredential.DistributionType == "APP_STORE" {
			appStoreMap = map[string]any{
				"id":                      buildCredential.Id,
				"certificate_id":          buildCredential.CertificateId,
				"provisioning_profile_id": buildCredential.ProvisioningProfileId,
			}
			if err := d.Set("app_store", []any{appStoreMap}); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}

	return diags
}
