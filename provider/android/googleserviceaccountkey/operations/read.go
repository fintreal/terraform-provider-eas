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

	identifier := d.Get("project_identifier").(string)
	input := eas.GetByProjectIdentifierGoogleServiceAccountKeyData{
		AccountId:         client.AccountId,
		ProjectIdentifier: identifier,
	}
	data, err := client.Android.GoogleServiceAccountKey.GetByProjectIdentifier(input)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)

	if err := d.Set("id", data.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("client_email", data.ClientEmail); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("project_identifier", data.ProjectIdentifier); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("client_identifier", data.ClientIdentifier); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
