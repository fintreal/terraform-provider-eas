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

	input := eas.CreateAppCredentialsData{
		AppId:           d.Get("app_id").(string),
		AppIdentifierId: d.Get("app_identifier_id").(string),
	}

	appStoreApiKeyId := d.Get("app_store_api_key_id").(string)
	if appStoreApiKeyId != "" {
		input.AppStoreApiKeyId = &appStoreApiKeyId
	}
	pushKeyId := d.Get("push_key_id").(string)
	if pushKeyId != "" {
		input.PushKeyId = &pushKeyId
	}

	data, err := client.Apple.AppCredentials.Create(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)
	d.Set("id", data.Id)
	d.Set("app_id", data.AppId)
	d.Set("app_identifier_id", data.AppIdentifierId)
	d.Set("app_store_api_key_id", data.AppStoreApiKeyId)
	d.Set("push_key_id", data.PushKeyId)

	appStoreList := d.Get("app_store").([]any)

	var diags diag.Diagnostics

	if len(appStoreList) > 0 {
		appStoreMap, ok := appStoreList[0].(map[string]any) // schema allows only 1 item
		if !ok {
			return diag.Errorf("Error while parsing app store config")
		}

		buildCredentialsInput := eas.CreateAppBuildCredentialsData{
			DistributionType:      "APP_STORE",
			CertificateId:         appStoreMap["certificate_id"].(string),
			ProvisioningProfileId: appStoreMap["provisioning_profile_id"].(string),
			AppCredentialsId:      data.Id,
		}
		buildData, err := client.Apple.AppBuildCredentials.Create(buildCredentialsInput)
		if err != nil {
			return diag.FromErr(err)
		}
		appStoreMap = map[string]any{
			"id":                      buildData.Id,
			"certificate_id":          buildData.CertificateId,
			"provisioning_profile_id": buildData.ProvisioningProfileId,
		}
		if err := d.Set("app_store", []any{appStoreMap}); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return diags
}
