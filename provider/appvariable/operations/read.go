package operations

import (
	"context"
	"terraform-provider-eas/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Read(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*client.EASClient)
	appId := d.Get("app_id").(string)
	name := d.Get("name").(string)

	data, err := client.AppVariable.GetByName(name, appId)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)

	if err := d.Set("id", data.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("name", data.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("value", data.Value); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("visibility", data.Visibility); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("environments", data.Environments); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("app_id", appId); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
