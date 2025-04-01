// // Copyright (c) HashiCorp, Inc.
// // SPDX-License-Identifier: MPL-2.0

package provider

// import (
// 	"context"
// 	"terraform-provider-expo-eas/internal/eas"

// 	"github.com/hashicorp/terraform-plugin-framework/path"
// 	"github.com/hashicorp/terraform-plugin-framework/resource"
// 	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
// 	"github.com/hashicorp/terraform-plugin-framework/types"
// )

// // Ensure provider defined types fully satisfy framework interfaces.
// var _ resource.Resource = &projectVariableResource{}
// var _ resource.ResourceWithImportState = &projectVariableResource{}

// func NewProjectVariableResource() resource.Resource {
// 	return &projectVariableResource{}
// }

// type projectVariableResource struct {
// }

// type ProjectVariableResourceModel struct {
// 	ProjectName types.String `tfsdk:"project_name"`
// 	Name        types.String `tfsdk:"name"`
// 	Value       types.String `tfsdk:"value"`
// 	Visibility  types.String `tfsdk:"visibility"`
// 	Environment types.String `tfsdk:"environment"`
// }

// func (r *projectVariableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
// 	resp.TypeName = req.ProviderTypeName + "_project_variable"
// }

// func (r *projectVariableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
// 	resp.Schema = schema.Schema{
// 		MarkdownDescription: "Project Resource",

// 		Attributes: map[string]schema.Attribute{
// 			"project_name": schema.StringAttribute{
// 				MarkdownDescription: "Project Name",
// 				Required:            true,
// 			},
// 			"name": schema.StringAttribute{
// 				MarkdownDescription: "Name",
// 				Required:            true,
// 			},
// 			"value": schema.StringAttribute{
// 				MarkdownDescription: "Value",
// 				Required:            true,
// 			},
// 			"visibility": schema.StringAttribute{
// 				MarkdownDescription: "Visibility",
// 				Required:            true,
// 			},
// 			"environment": schema.StringAttribute{
// 				MarkdownDescription: "Environment",
// 				Required:            true,
// 			},
// 		},
// 	}
// }

// func (r *projectVariableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
// }

// func (r *projectVariableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
// 	var data ProjectVariableResourceModel

// 	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	projectName := data.ProjectName.ValueString()
// 	projectVariableProps := eas.ProjectVariableProps{
// 		Name:        data.Name.ValueString(),
// 		Value:       data.Value.ValueString(),
// 		Environment: data.Environment.ValueString(),
// 		Visibility:  data.Visibility.ValueString(),
// 	}

// 	obj, err := eas.CreateProjectVariable(projectName, projectVariableProps)

// 	if err != nil {
// 		resp.Diagnostics.AddError(err.Error(), "")
// 		return
// 	}

// 	state := ProjectVariableResourceModel{
// 		Name:        types.StringValue(obj.Name),
// 		Value:       types.StringValue(obj.Value),
// 		Environment: types.StringValue(obj.Environment),
// 		Visibility:  types.StringValue(obj.Visibility),
// 		ProjectName: data.ProjectName,
// 	}

// 	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
// }

// func (r *projectVariableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
// 	var data ProjectVariableResourceModel

// 	// Read Terraform prior state data into the model
// 	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	out, err := eas.GetProjectVariable(
// 		data.ProjectName.ValueString(),
// 		data.Name.ValueString(),
// 		data.Environment.ValueString(),
// 	)
// 	if err != nil {
// 		resp.Diagnostics.AddError(err.Error(), "")
// 		return
// 	}
// 	data = ProjectVariableResourceModel{
// 		ProjectName: data.ProjectName,
// 		Name:        types.StringValue(out.Name),
// 		Environment: types.StringValue(out.Environment),
// 		Visibility:  types.StringValue(out.Visibility),
// 		Value:       types.StringValue(out.Value),
// 	}

// 	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
// }

// func (r *projectVariableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
// 	var data ProjectVariableResourceModel

// 	// Read Terraform plan data into the model
// 	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	// If applicable, this is a great opportunity to initialize any necessary
// 	// provider client data and make a call using it.
// 	// httpResp, err := r.client.Do(httpReq)
// 	// if err != nil {
// 	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update example, got error: %s", err))
// 	//     return
// 	// }

// 	// Save updated data into Terraform state
// 	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
// }

// func (r *projectVariableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
// 	var data ProjectVariableResourceModel

// 	// Read Terraform prior state data into the model
// 	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

// 	_, err := eas.DeleteProjectVariable(
// 		data.ProjectName.ValueString(),
// 		data.Name.ValueString(),
// 		data.Environment.ValueString(),
// 	)

// 	if err != nil {
// 		resp.Diagnostics.AddError(err.Error(), "")
// 	}
// }

// func (r *projectVariableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
// 	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
// }
