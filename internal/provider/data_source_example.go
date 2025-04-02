package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &exampleDataSource{}
var _ datasource.DataSourceWithConfigure = &exampleDataSource{}

func NewExampleDataSource() datasource.DataSource {
	return &exampleDataSource{}
}

type exampleDataSource struct {
	client *easClient
}

type exampleDataSourceModel struct {
	Id types.String `tfsdk:"id"`
}

func (d *exampleDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	// Change 'example'
	resp.TypeName = req.ProviderTypeName + "_example"
}

func (d *exampleDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Required: true,
		},
	}}
}

func (d *exampleDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, _ := req.ProviderData.(*easClient)

	d.client = client
}

func (d *exampleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config exampleDataSourceModel
	diags := resp.State.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	data, err := d.client.App.Get(config.Id.ValueString())

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
