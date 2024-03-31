package apiProperty

import (
	"context"
	"property-service/internal/model"
	"property-service/internal/model/enum"
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"
	"sync"
	"time"

	"github.com/hadanhtuan/go-sdk/common"
	"github.com/hadanhtuan/go-sdk/db/orm"
)

func (bc *PropertyController) CreateBooking(ctx context.Context, req *protoProperty.MsgBooking) (*protoSdk.BaseResponse, error) {
	if req.CheckInDate < time.Now().Unix() || req.CheckInDate >= req.CheckoutDate {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.BadRequest,
			Message: "Error booking. Checkin time not valid",
		})
	}

	var previousBooking model.Booking
	model.BookingDB.DB.Where("checkin_date BETWEEN ? OR ?", req.CheckInDate, req.CheckoutDate).
		Or("checkout_date BETWEEN ? OR ?", req.CheckInDate, req.CheckoutDate).
		First(&previousBooking)

	if previousBooking.ID != "" {
		return util.ConvertToGRPC(&common.APIResponse{
			Status:  common.APIStatus.ServerError,
			Message: "Error booking. Property already being booked",
		})
	}

	result := model.PropertyDB.QueryOne(&model.Property{
		ID: req.PropertyId,
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

func (bc *PropertyController) CountBookingStatus(ctx context.Context, req *protoProperty.MsgBooking) (*protoSdk.BaseResponse, error) {

	filter := &model.Booking{}
	statuses := util.ConvertEnumToSlice(*enum.BookingStatus)

	if req.Status != nil && *req.Status != "" {
		status := enum.BookingStatusValue(*req.Status)
		statuses = []string{string(status)}
	}

	if req.Id != "" {
		filter.ID = req.Id
	}

	if req.PropertyId != "" {
		filter.PropertyId = req.PropertyId
	}

	if req.UserId != "" {
		filter.UserId = req.UserId
	}

	result := make([]*model.CountByStatus, len(statuses))
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(statuses))

	for index, status := range statuses {
		go func(i int, s string, f model.Booking, r []*model.CountByStatus, wg *sync.WaitGroup) {
			defer wg.Done()

			status := enum.BookingStatusValue(s)
			f.Status = &status

			countResult := model.BookingDB.Count(f)

			r[i] = &model.CountByStatus{
				Status:   s,
				Quantity: &countResult.Total,
			}
		}(index, status, *filter, result, waitGroup)
	}
	waitGroup.Wait()

	return util.ConvertToGRPC(&common.APIResponse{
		Data:    result,
		Message: "Result for count by status",
		Status:  common.APIStatus.Ok,
	})
}

func (bc *PropertyController) GetBooking(ctx context.Context, req *protoProperty.MsgQueryBooking) (*protoSdk.BaseResponse, error) {
	queryField := req.QueryFields
	orderField := req.OrderFields

	filter := &model.Booking{}
	sort := []string{}

	if queryField.Id != "" {
		filter.ID = queryField.Id
	}

	if queryField.PropertyId != "" {
		filter.PropertyId = queryField.PropertyId
	}

	if queryField.UserId != "" {
		filter.UserId = queryField.UserId
	}

	if queryField.Status != nil && *queryField.Status != "" {
		status := enum.BookingStatusValue(*queryField.Status)
		filter.Status = &status
	}

	if orderField.CreatedAt != nil {
		sort = append(sort, "created_at "+orderField.CreatedAt.String())
	}

	result := model.BookingDB.Query(filter,
		req.Paginate.Offset,
		req.Paginate.Limit,
		&orm.QueryOption{
			Order: sort,
		},
	)

	return util.ConvertToGRPC(result)
}
