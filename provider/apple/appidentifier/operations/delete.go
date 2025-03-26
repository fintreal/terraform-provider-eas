package operations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Delete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	var warning = diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "App Identifier was removed from the state but not deleted!",
		Detail:   "EAS doesn't support the removal of app identifiers. The identifier still exists in your account. If you try to recreate the same identifier for the account, it will be imported into the state.",
	}
	return diag.Diagnostics{warning}
}
