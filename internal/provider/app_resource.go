// // Copyright (c) HashiCorp, Inc.
// // SPDX-License-Identifier: MPL-2.0

package provider

// import (
// 	"context"

// 	"github.com/hashicorp/terraform-plugin-framework/path"
// 	"github.com/hashicorp/terraform-plugin-framework/resource"
// 	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
// 	"github.com/hashicorp/terraform-plugin-framework/types"
// )

// // Ensure provider defined types fully satisfy framework interfaces.
// var _ resource.Resource = &ProjectResource{}
// var _ resource.ResourceWithImportState = &ProjectResource{}

// func NewProjectResource() resource.Resource {
// 	return &ProjectResource{}
// }

// type ProjectResource struct {
// }

// type projectResourceModel struct {
// 	Name  types.String `tfsdk:"name"`
// 	ID    types.String `tfsdk:"id"`
// 	Owner types.String `tfsdk:"owner"`
// }

// func (r *ProjectResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
// 	resp.TypeName = req.ProviderTypeName + "_project"
// }

// func (r *ProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
// 	resp.Schema = schema.Schema{
// 		MarkdownDescription: "Project Resource",

// 		Attributes: map[string]schema.Attribute{
// 			"name": schema.StringAttribute{
// 				MarkdownDescription: "Name",
// 				Required:            true,
// 			},
// 			"id": schema.StringAttribute{
// 				MarkdownDescription: "ID",
// 				Computed:            true,
// 			},
// 			"owner": schema.StringAttribute{
// 				MarkdownDescription: "Owner",
// 				Computed:            true,
// 			},
// 		},
// 	}
// }

// func (r *ProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

// }

// func (r *ProjectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
// 	var data projectResourceModel

// 	// Read Terraform plan data into the model
// 	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

// 	if resp.Diagnostics.HasError() {
// 		return
// 	}
// 	obj, err := eas.CreateProject(data.Name.ValueString())

// 	if err != nil {
// 		resp.Diagnostics.AddError(err.Error(), "")
// 		return
// 	}

// 	state := projectResourceModel{
// 		Name:  types.StringValue(obj.Name),
// 		Owner: types.StringValue(obj.Owner),
// 		ID:    types.StringValue(obj.ID),
// 	}

// 	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
// }

// func (r *ProjectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
// 	var data projectResourceModel

// 	// Read Terraform prior state data into the model
// 	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// 	// If applicable, this is a great opportunity to initialize any necessary
// 	// provider client data and make a call using it.
// 	// httpResp, err := r.client.Do(httpReq)
// 	// if err != nil {
// 	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
// 	//     return
// 	// }

// 	// Save updated data into Terraform state
// 	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
// }

// func (r *ProjectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
// 	var data projectResourceModel

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

// func (r *ProjectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
// 	var data projectResourceModel

// 	// Read Terraform prior state data into the model
// 	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

// 	if resp.Diagnostics.HasError() {
// 		return
// 	}

// }

// func (r *ProjectResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
// 	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
// }
