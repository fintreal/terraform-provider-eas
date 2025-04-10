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

	identifier := d.Get("identifier").(string)
	input := eas.CreateAppleAppIdentifierData{
		AccountId:  client.AccountId,
		Identifier: identifier,
	}

	data, err := client.Apple.AppIdentifier.Create(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)
	d.Set("id", data.Id)
	d.Set("identifier", data.Identifier)

	var diags diag.Diagnostics
	return diags
}
