package operations

import (
	"context"
	"terraform-provider-eas/internal/client"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Create(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*client.EASClient)

	input := eas.CreateAndroidAppCredentialsData{
		AppId:                     d.Get("app_id").(string),
		Identifier:                d.Get("identifier").(string),
		GoogleServiceAccountKeyId: d.Get("google_service_account_key_id").(string),
	}
	data, err := client.Android.AppCredentials.Create(input)

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
	if err := d.Set("identifier", data.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("google_service_account_key_id", data.GoogleServiceAccountKeyId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	buildCredentialsList := d.Get("build_credentials").([]any)

	buildCredentialsNewList := make([]any, len(buildCredentialsList))

	for _, buildCredentials := range buildCredentialsList {
		buildCredential, ok := buildCredentials.(map[string]any)
		if !ok {
			return diag.Errorf("Error while parsing build credentials config")
		}
		buildCredentialsInput := eas.CreateAndroidAppBuildCredentialsData{
			AppCredentialsId: data.Id,
			Name:             buildCredential["name"].(string),
			KeystoreId:       buildCredential["keystore_id"].(string),
		}

		buildData, err := client.Android.AppBuildCredentials.Create(buildCredentialsInput)
		if err != nil {
			return diag.FromErr(err)
		}
		buildCredentialsNewList = append(buildCredentialsNewList, map[string]any{
			"id":          buildData.Id,
			"name":        buildData.Name,
			"keystore_id": buildData.KeystoreId,
		})
	}

	if err := d.Set("build_credentials", buildCredentialsNewList); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
