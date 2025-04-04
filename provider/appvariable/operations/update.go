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

	value := d.Get("value").(string)
	visibility := d.Get("visibility").(string)

	set := d.Get("environments").(*schema.Set).List()

	var environments []string
	for _, v := range set {
		str := v.(string)
		environments = append(environments, str)
	}

	input := eas.UpdateAppVariableData{
		Id:           id,
		Name:         name,
		Value:        value,
		Visibility:   visibility,
		Environments: environments,
	}

	data, err := client.AppVariable.Update(input)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.Id)

	var diags diag.Diagnostics
	return diags
}
