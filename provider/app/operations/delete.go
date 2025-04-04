package operations

import (
	"context"
	"fmt"
	"terraform-provider-eas/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	client := m.(*client.EASClient)
	var warning = diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "App was removed from the state but not deleted!",
		Detail:   fmt.Sprintf("EAS requires elevated privilages to delete app. The app was removed from the state, but you have to delete it manually here https://expo.dev/accounts/%s/projects", client.AccountName),
	}
	return diag.Diagnostics{warning}
}
