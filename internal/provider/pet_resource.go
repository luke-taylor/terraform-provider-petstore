package provider

import (
    "context"
    "terraform-provider-petstore/internal/resource_pet"

    "github.com/hashicorp/terraform-plugin-framework/diag"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = (*petResource)(nil)

func NewPetResource() resource.Resource {
    return &petResource{}
}

type petResource struct{}

func (r *petResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
    resp.TypeName = req.ProviderTypeName + "_pet"
}

func (r *petResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
    resp.Schema = resource_pet.PetResourceSchema(ctx)
}

func (r *petResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
    var data resource_pet.PetModel

    resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(callPetAPI(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *petResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
    var data resource_pet.PetModel

    resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *petResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
    var data resource_pet.PetModel

    resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(callPetAPI(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *petResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
    var data resource_pet.PetModel

    resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
    if resp.Diagnostics.HasError() {
        return
    }
}

// Typically this method would contain logic that makes an HTTP call to a remote API, and then stores
// computed results back to the data model. For example purposes, this function just sets all unknown
// Pet values to null to avoid data consistency errors.
func callPetAPI(ctx context.Context, pet *resource_pet.PetModel) diag.Diagnostics {
    if pet.Id.IsUnknown() {
        pet.Id = types.Int64Null()
    }

    if pet.Status.IsUnknown() {
        pet.Status = types.StringNull()
    }

    if pet.Tags.IsUnknown() {
        pet.Tags = types.ListNull(resource_pet.TagsValue{}.Type(ctx))
    } else if !pet.Tags.IsNull() {
        var tags []resource_pet.TagsValue
        diags := pet.Tags.ElementsAs(ctx, &tags, false)
        if diags.HasError() {
            return diags
        }

        for i := range tags {
            if tags[i].Id.IsUnknown() {
                tags[i].Id = types.Int64Null()
            }

            if tags[i].Name.IsUnknown() {
                tags[i].Name = types.StringNull()
            }
        }

        pet.Tags, diags = types.ListValueFrom(ctx, resource_pet.TagsValue{}.Type(ctx), tags)
        if diags.HasError() {
            return diags
        }
    }

    if pet.Category.IsUnknown() {
        pet.Category = resource_pet.NewCategoryValueNull()
    } else if !pet.Category.IsNull() {
        if pet.Category.Id.IsUnknown() {
            pet.Category.Id = types.Int64Null()
        }

        if pet.Category.Name.IsUnknown() {
            pet.Category.Name = types.StringNull()
        }
    }

    return nil
}
