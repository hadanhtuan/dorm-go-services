package apiProperty

import (
	"context"
	"property-service/internal/model"
	"property-service/internal/model/enum"
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"
	"time"

	"github.com/hadanhtuan/go-sdk/common"
)

func (bc *PropertyController) CreateBooking(ctx context.Context, req *protoProperty.MsgBooking) (*protoSdk.BaseResponse, error) {
	if req.CheckInDate < time.Now().Unix() || req.CheckInDate >= req.CheckoutDate {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error booking. Checkin time not valid",
		})
	}

	var previousBooking model.Booking
	model.BookingDB.DB.Where("checkout_date >= ? OR checkin_date <= ?",
		req.CheckInDate,
		req.CheckoutDate,
	).First(&previousBooking)

	if previousBooking.ID != "" {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.ServerError,
			Message: "Error booking. Property already being booked",
		})
	}



	result := model.PropertyDB.QueryOne(map[string]interface{}{
		"id": req.PropertyId,
	}, nil)

	if result.Status == common.APIStatus.NotFound {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.ServerError,
			Message: "Error booking. Not found property",
		})
	}
	property := result.Data.([]*model.Property)[0]

	totalPriceBeforeTax := float64(req.NightNumber)*property.NightPrice + property.ServiceFee
	taxFee := totalPriceBeforeTax * property.TaxPercent
	totalPrice := totalPriceBeforeTax + taxFee

	booking := &model.Booking{
		PropertyId:          req.PropertyId,
		UserId:              req.UserId,
		CheckInDate:         req.CheckInDate,
		CheckoutDate:        req.CheckoutDate,
		GuestNumber:         req.GuestNumber,
		ChildNumber:         req.ChildNumber,
		BabyNumber:          req.BabyNumber,
		PetNumber:           req.PetNumber,
		NightNumber:         req.NightNumber,
		TotalPriceBeforeTax: totalPriceBeforeTax,
		TotalPrice:          totalPrice,
		TaxFee:              taxFee,
		Status:              &enum.BookingStatus.Created,
	}

	result = model.BookingDB.Create(booking)

	return util.ConvertToGRPC(result)
}
