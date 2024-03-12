package apiBooking

import (
	"booking-service/internal/model"
	"booking-service/internal/model/enum"
	"booking-service/internal/util"
	protoBooking "booking-service/proto/booking"
	protoSdk "booking-service/proto/sdk"
	"context"
	"encoding/json"

	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/common"
)

func (bc *BookingController) GetPropertyDetail(ctx context.Context, req *protoBooking.MsgGetProperty) (*protoSdk.BaseResponse, error) {
	property := &model.Property{
		ID: req.PropertyId,
	}

	result := model.PropertyDB.QueryOne(property)
	if result.Status == common.APIStatus.NotFound {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.NotFound,
			Message: "Property Not Found",
		})
	}
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) GetAllProperty(ctx context.Context, req *protoBooking.MsgQueryProperty) (*protoSdk.BaseResponse, error) {
	result := model.PropertyDB.Query(nil, req.Paginate.Offset, req.Paginate.Limit)

	result.Message = "Get all properties successfully"
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) CreateProperty(ctx context.Context, req *protoBooking.MsgCreateProperty) (*protoSdk.BaseResponse, error) {
	propertyType := enum.PropertyTypeValue(req.PropertyType)

	property := &model.Property{
		HostId:       req.HostId,
		WardId:       req.WardId,
		NumBeds:      req.NumBed,
		NumBedrooms:  req.NumBedroom,
		NumBaths:     req.NumBath,
		IsGuestFavor: req.IsGuestFavor,
		Body:         req.Body,
		Title:        req.Title,
		// Amenities:    strings.Join(req.Amenities, "|"),
		PropertyType: &propertyType,
	}

	result := model.PropertyDB.Create(property)
	if result.Status != common.APIStatus.Created {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.ServerError,
			Message: "Create Property Failed.",
		})
	}

	//go routine
	data := result.Data.([]*model.Property)[0]
	
	go bc.SyncProperty(data)

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Create Property Successfully.",
	})
}

func (bc *BookingController) UpdateProperty(ctx context.Context, req *protoBooking.MsgUpdateProperty) (*protoSdk.BaseResponse, error) {
	propertyId := req.PropertyId
	property := &model.Property{
		ID: propertyId,
	}
	propertyType := enum.PropertyTypeValue(req.PropertyType)

	propertyUpdated := &model.Property{
		NumBeds:      req.NumBed,
		NumBedrooms:  req.NumBedroom,
		NumBaths:     req.NumBath,
		IsGuestFavor: req.IsGuestFavor,
		Body:         req.Body,
		Title:        req.Title,
		// Amenities:    strings.Join(req.Amenities, "|"),
		PropertyType: &propertyType,
	}

	result := model.PropertyDB.Update(property, propertyUpdated)
	return util.ConvertToGRPC(result)

}

func (bc *BookingController) DeleteProperty(ctx context.Context, req *protoBooking.MsgDeleteProperty) (*protoSdk.BaseResponse, error) {
	propertyId := req.PropertyId
	property := &model.Property{
		ID: propertyId,
	}

	result := model.PropertyDB.Delete(property)
	return util.ConvertToGRPC(result)
}
