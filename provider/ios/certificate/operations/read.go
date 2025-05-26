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

	input := eas.GetBySerialNumberAppleCertificateData{
		AccountId:    client.AccountId,
		SerialNumber: d.Get("serial_number").(string),
	}
	data, err := client.Apple.Certificate.GetBySerialNumber(input)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Id)

	if err := d.Set("id", data.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("serial_number", data.SerialNumber); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}
