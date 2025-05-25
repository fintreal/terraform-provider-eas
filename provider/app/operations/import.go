package operations

import (
	"context"
	"terraform-provider-eas/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Import(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	client := m.(*client.EASClient)

	data, err := client.App.Get(d.Id())
	if err != nil {
		return nil, err
	}

	if err := d.Set("id", data.Id); err != nil {
		return nil, err
	}
	if err := d.Set("name", data.Name); err != nil {
		return nil, err
	}
	if err := d.Set("slug", data.Slug); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
