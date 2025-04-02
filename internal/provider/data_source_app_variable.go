package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &appVariableDataSource{}

func NewAppVariableDataSource() datasource.DataSource {
	return &appVariableDataSource{}
}

type appVariableDataSource struct {
	client easClient
}

type appVariableDataSourceModel struct {
	AppId        types.String `tfsdk:"app_id"`
	Id           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Value        types.String `tfsdk:"value"`
	Visibility   types.String `tfsdk:"visibility"`
	Environments types.Set    `tfsdk:"environments"`
}

func (d *appVariableDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_variable"
}

func (d *appVariableDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"app_id": schema.StringAttribute{
			Required: true,
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"id": schema.StringAttribute{
			Computed: true,
		},
		"value": schema.StringAttribute{
			Computed: true,
		},
		"visibility": schema.StringAttribute{
			Computed:   true,
			Validators: []validator.String{}, // TODO
		},
		"environments": schema.SetAttribute{
			ElementType: types.StringType,
			Computed:    true,
		},
	}}
}

func (d *appVariableDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, _ := req.ProviderData.(*easClient)

	d.client = *client
}

func (d *appVariableDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config appVariableDataSourceModel
	diags := resp.State.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data, err := d.client.AppEnvVar.GetByName(config.Name.ValueString(), config.AppId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Unable to Read 'app'", err.Error())
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
		Visibility:   types.StringValue(string(data.Visibility)),
		Environments: types.SetValueMust(types.StringType, environments),
	}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
