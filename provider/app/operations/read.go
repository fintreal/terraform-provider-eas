package operations

import (
	"context"
	"terraform-provider-eas/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)
	id := d.Get("id").(string)

	d.SetId(id)
	data, err := client.App.Get(id)

	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics
	if err := d.Set("name", data.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("slug", data.Slug); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
