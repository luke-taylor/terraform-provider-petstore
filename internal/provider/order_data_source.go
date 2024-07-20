package provider

import (
    "context"
    "terraform-provider-petstore/internal/datasource_order"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/diag"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*orderDataSource)(nil)

func NewOrderDataSource() datasource.DataSource {
    return &orderDataSource{}
}

type orderDataSource struct{}

func (d *orderDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_order"
}

func (d *orderDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
    resp.Schema = datasource_order.OrderDataSourceSchema(ctx)
}

func (d *orderDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
    var data datasource_order.OrderModel

    resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(callOrderAPI(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Typically this method would contain logic that makes an HTTP call to a remote API, and then stores
// computed results back to the data model. For example purposes, this function just sets computed Order
// values to mock values to avoid data consistency errors.
func callOrderAPI(ctx context.Context, order *datasource_order.OrderModel) diag.Diagnostics {
    order.PetId = types.Int64Value(1)
    order.Quantity = types.Int64Value(10)
    order.Status = types.StringValue("delivered")
    order.ShipDate = types.StringValue("2023-07-25T23:43:16Z")
    order.Complete = types.BoolValue(true)

    return nil
}
