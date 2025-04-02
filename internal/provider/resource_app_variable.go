package provider

import (
	"context"

	"github.com/fintreal/expo-eas-sdk-go/eas"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &appVariableResource{}
var _ resource.ResourceWithConfigure = &appVariableResource{}

func NewAppVariableResource() resource.Resource {
	return &appVariableResource{}
}

type appVariableResource struct {
	client *easClient
}

type appVariableResourceModel struct {
	AppId        types.String `tfsdk:"app_id"`
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Value        types.String `tfsdk:"value"`
	Visibility   types.String `tfsdk:"visibility"`
	Environments types.Set    `tfsdk:"environments"`
}

func (r *appVariableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_variable"
}

func (d *appVariableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, _ := req.ProviderData.(*easClient)

	d.client = client
}

func (r *appVariableResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"app_id": schema.StringAttribute{
			Required: true,
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"value": schema.StringAttribute{
			Required: true,
		},
		"visibility": schema.StringAttribute{
			Required:   true,
			Validators: []validator.String{}, // TODO
		},
		"environments": schema.SetAttribute{
			ElementType: types.StringType,
			Required:    true,
		},
	}}
}

func (r *appVariableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var config appVariableResourceModel
	diags := resp.State.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data, err := r.client.AppEnvVar.Get(config.Id.ValueString(), config.AppId.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Unable to Read 'app_variable'", err.Error())
		return
	}

	var environments []attr.Value
	for _, environment := range data.Environments {
		environments = append(environments, types.StringValue(string(environment)))
	}
	state := appVariableDataSourceModel{
		Id:           types.StringValue(data.Id),
		AppId:        types.StringValue(data.AppId),
		Name:         types.StringValue(data.Name),
		Value:        types.StringValue(data.Value),
		Visibility:   types.StringValue(data.Visibility),
		Environments: types.SetValueMust(types.StringType, environments),
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Create creates the resource and sets the initial Terraform state.
func (r *appVariableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan appVariableResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var easEnvironments []string
	for _, environment := range plan.Environments.Elements() {
		easEnvironments = append(easEnvironments, (environment.(types.String)).ValueString())
	}

	appVariable := eas.CreateAppEnvVarData{
		AppId:        plan.AppId.ValueString(),
		Name:         plan.Name.ValueString(),
		Value:        plan.Value.ValueString(),
		Visibility:   plan.Visibility.ValueString(),
		Environments: easEnvironments,
	}
	data, err := r.client.AppEnvVar.Create(appVariable)

	if err != nil {
		resp.Diagnostics.AddError("Unable to Create'app_variable'", err.Error())
		return
	}

	var environments []attr.Value
	for _, environment := range data.Environments {
		environments = append(environments, types.StringValue(environment))
	}
	state := appVariableResourceModel{
		Id:           types.StringValue(data.Id),
		AppId:        types.StringValue(data.AppId),
		Name:         types.StringValue(data.Name),
		Value:        types.StringValue(data.Value),
		Visibility:   types.StringValue(data.Visibility),
		Environments: types.SetValueMust(types.StringType, environments),
	}
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *appVariableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *appVariableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
