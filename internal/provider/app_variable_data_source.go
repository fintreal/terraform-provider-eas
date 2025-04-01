// // Copyright (c) HashiCorp, Inc.
// // SPDX-License-Identifier: MPL-2.0

package provider

// import (
// 	"context"
// 	"terraform-provider-expo-eas/internal/eas"

// 	"github.com/hashicorp/terraform-plugin-framework/datasource"
// 	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
// 	"github.com/hashicorp/terraform-plugin-framework/types"
// )

// func NewAppVariableDataSource() datasource.DataSource {
// 	return &appVariableDataSource{}
// }

// type appVariableDataSource struct {
// }

// type appVariableDataSourceModel struct {
// 	Name        types.String `tfsdk:"name"`
// 	Environment types.String `tfsdk:"environment"`
// 	Value       types.String `tfsdk:"value"`
// 	Visibility  types.String `tfsdk:"visibility"`
// 	ProjectName types.String `tfsdk:"project_name"`
// }

// func (d *appVariableDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
// 	resp.TypeName = req.ProviderTypeName + "_project_variable"
// }

// func (d *appVariableDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
// 	resp.Schema = schema.Schema{
// 		MarkdownDescription: "Project Variable Data Source",

// 		Attributes: map[string]schema.Attribute{
// 			"project_name": schema.StringAttribute{
// 				MarkdownDescription: "Project name for the variable",
// 				Required:            true,
// 			},
// 			"name": schema.StringAttribute{
// 				MarkdownDescription: "Name for the variable",
// 				Required:            true,
// 			},
// 			"environment": schema.StringAttribute{
// 				MarkdownDescription: "Environment for the variable",
// 				Required:            true,
// 			},
// 			"visibility": schema.StringAttribute{
// 				MarkdownDescription: "Visibility",
// 				Computed:            true,
// 			},
// 			"value": schema.StringAttribute{
// 				MarkdownDescription: "Value",
// 				Computed:            true,
// 			},
// 		},
// 	}
// }

// func (d *appVariableDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
// 	var current appVariableDataSourceModel
// 	diags := resp.State.Get(ctx, &current)
// 	resp.Diagnostics.Append(diags...)
// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	var state appVariableDataSourceModel
// 	//fmt.Println(getProjectVariable(ctx.projectName, ctx.variableName, ctx.environment))
// 	obj, err := eas.GetProjectVariable(
// 		current.ProjectName.ValueString(),
// 		current.Name.ValueString(),
// 		current.Environment.ValueString(),
// 	)
// 	if err != nil {
// 		resp.Diagnostics.AddError(err.Error(), "")
// 		return
// 	}
// 	state = appVariableDataSourceModel{
// 		Name:        types.StringValue(obj.Name),
// 		ProjectName: current.ProjectName,
// 		Value:       types.StringValue(obj.Value),
// 		Environment: types.StringValue(obj.Environment),
// 		Visibility:  types.StringValue(obj.Visibility),
// 	}
// 	diags = resp.State.Set(ctx, &state)
// 	resp.Diagnostics.Append(diags...)
// }
