// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_order

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OrderDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"complete": schema.BoolAttribute{
				Computed: true,
			},
			"id": schema.Int64Attribute{
				Required:            true,
				Description:         "ID of order that needs to be fetched",
				MarkdownDescription: "ID of order that needs to be fetched",
			},
			"pet_id": schema.Int64Attribute{
				Computed: true,
			},
			"quantity": schema.Int64Attribute{
				Computed: true,
			},
			"ship_date": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed:            true,
				Description:         "Order Status",
				MarkdownDescription: "Order Status",
			},
		},
	}
}

type OrderModel struct {
	Complete types.Bool   `tfsdk:"complete"`
	Id       types.Int64  `tfsdk:"id"`
	PetId    types.Int64  `tfsdk:"pet_id"`
	Quantity types.Int64  `tfsdk:"quantity"`
	ShipDate types.String `tfsdk:"ship_date"`
	Status   types.String `tfsdk:"status"`
}