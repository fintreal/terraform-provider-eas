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

	input := eas.GetAndroidAppCredentialsData{
		Id:    d.Get("id").(string),
		AppId: d.Get("app_id").(string),
	}

	data, err := client.Android.AppCredentials.Get(input)

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
	if err := d.Set("identifier", data.Identifier); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("google_service_account_key_id", data.GoogleServiceAccountKeyId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("fcm_key", data.FCMKey); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	buildCredentials := []map[string]any{}
	for _, buildCredential := range data.BuildCredentials {
		buildCredentialMap := map[string]any{
			"id":          buildCredential.Id,
			"name":        buildCredential.Name,
			"keystore_id": buildCredential.KeystoreId,
		}
		buildCredentials = append(buildCredentials, buildCredentialMap)
	}

	if err := d.Set("build_credentials", buildCredentials); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
