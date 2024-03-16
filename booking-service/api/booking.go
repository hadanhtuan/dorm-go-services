package apiBooking

import (
	"booking-service/internal/model"
	"booking-service/internal/model/enum"
	"booking-service/internal/util"
	protoBooking "booking-service/proto/booking"
	protoSdk "booking-service/proto/sdk"
	"context"
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/hadanhtuan/go-sdk/db/orm"
)

// Booking
func (bc *BookingController) GetBookingDetail(ctx context.Context, req *protoBooking.MsgGetBooking) (*protoSdk.BaseResponse, error) {

	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Get Booking Detail Successfully.",
	})
}

func (bc *BookingController) GetProperty(ctx context.Context, req *protoBooking.MsgQueryProperty) (*protoSdk.BaseResponse, error) {
	filter := &model.Property{
		Reviews: []*model.Review{},
	}

	if req.QueryFields.Id != nil {
		filter.ID = *req.QueryFields.Id
	}

	result := model.PropertyDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, &orm.QueryOption{
		Preload: []string{"Reviews"}, //Field name, not table name
	})

	result.Message = "Get properties successfully"
	return util.ConvertToGRPC(result)
}

func (bc *BookingController) CreateProperty(ctx context.Context, req *protoBooking.MsgProperty) (*protoSdk.BaseResponse, error) {
	propertyType := enum.PropertyTypeValue(req.PropertyType)
	property := &model.Property{
		HostId:       req.HostId,
		WardCode:     req.WardCode,
		NationCode:   req.NationCode,
		NumBeds:      req.NumBeds,
		NumBedrooms:  req.NumBedrooms,
		NumBathrooms: req.NumBathrooms,
		IsGuestFavor: req.IsGuestFavor,
		Body:         req.Body,
		Title:        req.Title,
		PropertyType: &propertyType,
	}

	result := model.PropertyDB.Create(property)
	data := result.Data.([]*model.Property)[0]

	// Sync data to search service
	go bc.SyncProperty(data)

	if result.Status != common.APIStatus.Created {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.ServerError,
			Message: "Create Property Failed.",
		})
	}
	return util.ConvertToGRPC(&common.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Create Property Successfully.",
	})
}

func (bc *BookingController) UpdateProperty(ctx context.Context, req *protoBooking.MsgProperty) (*protoSdk.BaseResponse, error) {
	property := &model.Property{
		ID: *req.Id,
	}
	propertyType := enum.PropertyTypeValue(req.PropertyType)

	propertyUpdated := &model.Property{
		NumBeds:      req.NumBeds,
		NumBedrooms:  req.NumBedrooms,
		NumBathrooms: req.NumBathrooms,
		IsGuestFavor: req.IsGuestFavor,
		Body:         req.Body,
		Title:        req.Title,
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
