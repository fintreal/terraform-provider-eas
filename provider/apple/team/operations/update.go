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

	id := d.Id()
	name := d.Get("name").(string)
	teamType := d.Get("type").(string)

	input := eas.UpdateAppleTeamData{
		Id:   id,
		Name: name,
		Type: teamType,
	}

	data, err := client.Apple.Team.Update(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)

	var diags diag.Diagnostics
	return diags
}
