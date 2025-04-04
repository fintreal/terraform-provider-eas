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
	appId := d.Get("app_id").(string)
	value := d.Get("value").(string)
	visibility := d.Get("visibility").(string)

	set := d.Get("environments").(*schema.Set).List()

	var environments []string
	for _, v := range set {
		str := v.(string)
		environments = append(environments, str)
	}

	input := eas.CreateAppVariableData{
		Name:         name,
		AppId:        appId,
		Value:        value,
		Visibility:   visibility,
		Environments: environments,
	}

	data, err := client.AppVariable.Create(input)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.Id)

	var diags diag.Diagnostics
	return diags
}
