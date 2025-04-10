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

	data, err := client.Apple.AppCredentials.Create(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)
	d.Set("id", data.Id)
	d.Set("app_id", data.AppId)
	d.Set("app_identifier_id", data.AppIdentifierId)

	var diags diag.Diagnostics
	return diags
}
