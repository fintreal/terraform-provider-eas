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

// func handleCredentialsChange(d *schema.ResourceData, client *client.EASClient) diag.Diagnostics {
// 	var diags diag.Diagnostics

// 	if !d.HasChange("app_store_api_key_id") && !d.HasChange("push_key_id") {
// 		return diags
// 	}

// 	input := eas.UpdateAppCredentialsData{
// 		Id: d.Get("id").(string),
// 	}

// 	appStoreApiKeyId := d.Get("app_store_api_key_id").(string)
// 	if appStoreApiKeyId != "" {
// 		input.AppStoreApiKeyId = &appStoreApiKeyId
// 	}
// 	pushKeyId := d.Get("push_key_id").(string)
// 	if pushKeyId != "" {
// 		input.PushKeyId = &pushKeyId
// 	}

// 	diags = append(diags, diag.Errorf("%+v\n", input)...)

// 	data, err := client.Apple.AppCredentials.Update(input)
// 	diags = append(diags, diag.Errorf("%+v\n", *data.AppStoreApiKeyId)...)

// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	d.SetId(data.Id)
// 	d.Set("id", data.Id)
// 	d.Set("app_id", data.AppId)
// 	d.Set("app_identifier_id", data.AppIdentifierId)
// 	d.Set("app_store_api_key_id", data.AppStoreApiKeyId)
// 	d.Set("push_key_id", data.PushKeyId)

// 	return diags
// }

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
