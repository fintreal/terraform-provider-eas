package operations

import (
	"context"
	"terraform-provider-eas/internal/client"

	"github.com/fintreal/eas-sdk-go/eas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Update(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	if !d.HasChange("build_credentials") {
		return diags
	}

	client := m.(*client.EASClient)

	old, new := d.GetChange("build_credentials")
	oldList := old.([]any)
	newList := new.([]any)
	for _, b := range oldList {
		buildCredential := b.(map[string]any)
		id := buildCredential["id"].(string)
		_, err := client.Android.AppBuildCredentials.Delete(id)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	for _, b := range newList {
		buildCredential := b.(map[string]any)
		buildCredentialsInput := eas.CreateAndroidAppBuildCredentialsData{
			AppCredentialsId: d.Get("id").(string),
			Name:             buildCredential["name"].(string),
			KeystoreId:       buildCredential["keystore_id"].(string),
		}

		data, err := client.Android.AppBuildCredentials.Create(buildCredentialsInput)
		if err != nil {
			return diag.FromErr(err)
		}
		newList = append(newList, map[string]any{
			"id":          data.AppCredentialsId,
			"keystore_id": data.KeystoreId,
			"name":        data.Name,
		})
	}
	if err := d.Set("build_credentials", newList); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}
