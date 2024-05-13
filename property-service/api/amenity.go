package apiProperty

import (
	"property-service/internal/model"
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"
	"context"
)

func (bc *PropertyController) CreateAmenity(ctx context.Context, req *protoProperty.MsgAmenity) (*protoSdk.BaseResponse, error) {
	amenity := &model.Amenity{
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	}

	result := model.AmenityDB.Create(amenity)

	return util.ConvertToGRPC(result)
}

func (bc *PropertyController) UpdateAmenity(ctx context.Context, req *protoProperty.MsgAmenity) (*protoSdk.BaseResponse, error) {
	amenity := &model.Amenity{
		ID: req.Id,
	}

	amenityUpdated := &model.Amenity{
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	}

	result := model.AmenityDB.Update(amenity, amenityUpdated)
	return util.ConvertToGRPC(result)

}

func (bc *PropertyController) DeleteAmenity(ctx context.Context, req *protoProperty.MsgId) (*protoSdk.BaseResponse, error) {
	amenity := &model.Amenity{
		ID: req.Id,
	}

	result := model.AmenityDB.Delete(amenity)
	return util.ConvertToGRPC(result)
}

func (bc *PropertyController) GetAmenity(ctx context.Context, req *protoProperty.MsgQueryAmenity) (*protoSdk.BaseResponse, error) {
	filter := &model.Amenity{}

	if req.QueryFields.Id != "" {
		filter.ID = req.QueryFields.Id
	}

	if req.QueryFields.Name != "" {
		filter.Name = req.QueryFields.Name
	}

	result := model.AmenityDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, nil)

	return util.ConvertToGRPC(result)
}
