package provider

import (
	"context"

	"github.com/fintreal/expo-eas-sdk-go/eas"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &appResource{}
var _ resource.ResourceWithConfigure = &appResource{}

func NewAppResource() resource.Resource {
	return &appResource{}
}

type appResource struct {
	client *easClient
}

type appResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Slug types.String `tfsdk:"slug"`
}

func (r *appResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

func (d *appResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, _ := req.ProviderData.(*easClient)

	d.client = client
}

func (r *appResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"slug": schema.StringAttribute{
			Required: true,
		},
	}}
}

func (r *appResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var config appResourceModel
	diags := resp.State.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data, err := r.client.App.Get(config.Id.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Unable to Read 'app'", err.Error())
		return
	}

	state := appResourceModel{
		Id:   types.StringValue(data.Id),
		Name: types.StringValue(data.Name),
		Slug: types.StringValue(data.Slug),
	}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Create creates the resource and sets the initial Terraform state.
func (r *appResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan appResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := eas.CreateAppData{
		Name:      plan.Name.ValueString(),
		Slug:      plan.Slug.ValueString(),
		AccountId: r.client.accountId,
	}

	data, err := r.client.App.Create(input)

	if err != nil {
		resp.Diagnostics.AddError("Unable to Create'app'", err.Error())
		return
	}
	state := appResourceModel{
		Id:   types.StringValue(data.Id),
		Name: types.StringValue(data.Name),
		Slug: types.StringValue(data.Slug),
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *appResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan appResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var currentState appResourceModel
	diags = req.State.Get(ctx, &currentState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not possible to update slug!
	if plan.Slug.ValueString() != currentState.Slug.ValueString() {
		resp.Diagnostics.AddAttributeError(path.Root("slug"), "Error Updating 'app'", "App slug cannot be changed!")
	}
	if resp.Diagnostics.HasError() {
		return
	}

	input := eas.UpdateAppData{
		Id:   currentState.Id.ValueString(),
		Name: plan.Name.ValueString(),
	}

	data, err := r.client.App.Update(input)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating 'app'", err.Error())
		return
	}

	newState := appResourceModel{
		Id:   types.StringValue(data.Id),
		Name: types.StringValue(data.Name),
		Slug: types.StringValue(data.Slug),
	}

	diags = resp.State.Set(ctx, newState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *appResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
