package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &exampleResource{}
var _ resource.ResourceWithConfigure = &exampleResource{}

func NewExampleResource() resource.Resource {
	return &exampleResource{}
}

type exampleResource struct {
	client *easClient
}

func (r *exampleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	// Change 'example'
	resp.TypeName = req.ProviderTypeName + "_example"
}

func (d *exampleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, _ := req.ProviderData.(*easClient)

	d.client = client
}

func (r *exampleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Required: true,
		},
	}}
}

// Create creates the resource and sets the initial Terraform state.
func (r *exampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read refreshes the Terraform state with the latest data.
func (r *exampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var config exampleDataSourceModel
	diags := resp.State.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data, err := r.client.App.Get(config.Id.ValueString())

	if err != nil {
		// Change 'example'
		resp.Diagnostics.AddError("Unable to Read 'example'", err.Error())
		return
	}

	state := exampleDataSourceModel{
		Id: types.StringValue(data.Id),
	}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *exampleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *exampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
