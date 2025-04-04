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

	name := d.Get("name").(string)
	slug := d.Get("slug").(string)

	input := eas.CreateAppData{
		Name:      name,
		Slug:      slug,
		AccountId: client.AccountId,
	}

	data, err := client.App.Create(input)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.Id)

	var diags diag.Diagnostics
	return diags
}
