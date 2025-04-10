package operations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var warning = diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Apple Team was removed from the state but not deleted!",
		Detail:   "The team was removed from the state, but it cannot be deleted.",
	}
	return diag.Diagnostics{warning}
}
