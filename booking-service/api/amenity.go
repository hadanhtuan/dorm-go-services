package apiBooking

import (
	"booking-service/internal/model"
	"booking-service/internal/util"
	protoBooking "booking-service/proto/booking"
	protoSdk "booking-service/proto/sdk"
	"context"
)

func (bc *BookingController) CreateAmenity(ctx context.Context, req *protoBooking.MsgAmenity) (*protoSdk.BaseResponse, error) {
	amenity := &model.Amenity{
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	}

	result := model.AmenityDB.Create(amenity)

	return util.ConvertToGRPC(result)
}

func (bc *BookingController) UpdateAmenity(ctx context.Context, req *protoBooking.MsgAmenity) (*protoSdk.BaseResponse, error) {
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

func (bc *BookingController) DeleteAmenity(ctx context.Context, req *protoBooking.MsgId) (*protoSdk.BaseResponse, error) {
	amenity := &model.Amenity{
		ID: req.Id,
	}

	result := model.AmenityDB.Delete(amenity)
	return util.ConvertToGRPC(result)
}

func (bc *BookingController) GetAmenity(ctx context.Context, req *protoBooking.MessageQueryAmenity) (*protoSdk.BaseResponse, error) {
	filter := &model.Amenity{}

	if req.QueryFields.Id != "" {
		filter.ID = req.QueryFields.Id
	}

	result := model.AmenityDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, nil)

	return util.ConvertToGRPC(result)
}
