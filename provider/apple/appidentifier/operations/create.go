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

	getInput := eas.GetByIdentifierAppleAppIdentifierData{
		AccountId:  client.AccountId,
		Identifier: d.Get("identifier").(string),
	}
	getData, err := client.Apple.AppIdentifier.GetByIdentifier(getInput)

	if err == nil {
		d.SetId(getData.Id)
		d.Set("id", getData.Id)
		d.Set("identifier", getData.Identifier)
		return diag.Diagnostics{diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "App Identifier " + identifier + " already exists. Importing it into the state!",
		}}
	}

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
