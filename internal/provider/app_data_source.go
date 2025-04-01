package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &appDataSource{}

func NewAppDataSource() datasource.DataSource {
	return &appDataSource{}
}

type appDataSource struct {
	client easClient
}

type appDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Slug types.String `tfsdk:"slug"`
}

func (d *appDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app"
}

func (d *appDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Required: true,
		},
		"name": schema.StringAttribute{
			Computed: true,
		},
		"slug": schema.StringAttribute{
			Computed: true,
		},
	}}
}

func (d *appDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, _ := req.ProviderData.(*easClient)

	d.client = *client
}

func (d *appDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config appDataSourceModel
	diags := resp.State.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data, err := d.client.App.Get(config.Id.ValueString())

	if err != nil {
		resp.Diagnostics.AddError("Unable to Read 'app'", err.Error())
		return
	}

	state := appDataSourceModel{
		Id:   types.StringValue(data.Id),
		Name: types.StringValue(data.Name),
		Slug: types.StringValue(data.Slug),
	}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
