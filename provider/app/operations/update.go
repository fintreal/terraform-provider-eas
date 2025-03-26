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

	id := d.Get("id").(string)
	name := d.Get("name").(string)

	d.SetId(id)
	input := eas.UpdateAppData{
		Id:   id,
		Name: name,
	}

	data, err := client.App.Update(input)
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics
	if err := d.Set("id", data.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("name", data.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("slug", data.Slug); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
