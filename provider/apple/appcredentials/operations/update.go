package operations

import (
	"context"
	"terraform-provider-eas/internal/client"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)

	var diags diag.Diagnostics
	diags = append(diags, handleAppStoreChange(d, client)...)
	//diags = append(diags, handleCredentialsChange(d, client)...)
	return diags
}

func handleAppStoreChange(d *schema.ResourceData, client *client.EASClient) diag.Diagnostics {
	var diags diag.Diagnostics
	if !d.HasChange("app_store") {
		return diags
	}
	old, new := d.GetChange("app_store")
	oldList := old.([]any)
	newList := new.([]any)
	if len(oldList) > 0 {
		oldMap := oldList[0].(map[string]any)
		id := oldMap["id"].(string)
		_, err := client.Apple.AppBuildCredentials.Delete(id)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if len(newList) > 0 {
		newMap := newList[0].(map[string]any)
		buildCredentialsInput := eas.CreateAppBuildCredentialsData{
			DistributionType:      "APP_STORE",
			CertificateId:         newMap["certificate_id"].(string),
			ProvisioningProfileId: newMap["provisioning_profile_id"].(string),
			AppCredentialsId:      d.Get("id").(string),
		}
		buildData, err := client.Apple.AppBuildCredentials.Create(buildCredentialsInput)
		if err != nil {
			return diag.FromErr(err)
		}
		newMap = map[string]any{
			"id":                      buildData.Id,
			"certificate_id":          buildData.CertificateId,
			"provisioning_profile_id": buildData.ProvisioningProfileId,
		}
		if err := d.Set("app_store", []any{newMap}); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}
