package apiProperty

import (
	"context"
	"property-service/internal/model"
	"property-service/internal/model/enum"
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"
	"sync"

	"github.com/hadanhtuan/go-sdk/common"
	"github.com/hadanhtuan/go-sdk/db/orm"
)

func (bc *PropertyController) GetProperty(ctx context.Context, req *protoProperty.MsgQueryProperty) (*protoSdk.BaseResponse, error) {
	filter := &model.Property{}

	if req.QueryFields.Id != nil {
		filter.ID = *req.QueryFields.Id
	}

	result := model.PropertyDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, &orm.QueryOption{
		Preload: []string{"Reviews", "Amenities", "Bookings"}, //Field name, not table name
		Order:   []string{"created_at desc"},
	})

	result.Message = "Get properties successfully"
	return util.ConvertToGRPC(result)
}

func (bc *PropertyController) CountPropertyStatus(ctx context.Context, req *protoProperty.MsgProperty) (*protoSdk.BaseResponse, error) {

	filter := &model.Property{}
	statuses := util.ConvertEnumToSlice(*enum.PropertyStatus)

	if req.Status != nil && *req.Status != "" {
		status := enum.BookingStatusValue(*req.Status)
		statuses = []string{string(status)}
	}

	if req.Id != nil {
		filter.ID = *req.Id
	}

	if req.HostId != "" {
		filter.HostId = req.HostId
	}

	result := make([]*model.CountByStatus, len(statuses))
	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(len(statuses))

	for index, status := range statuses {
		// hint: user slice instead of channel
		go func(i int, s string, f model.Property, r []*model.CountByStatus, wg *sync.WaitGroup) {
			defer wg.Done()

			status := enum.PropertyStatusValue(s)
			f.Status = &status

			countResult := model.PropertyDB.Count(f)

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

func (bc *PropertyController) CreateProperty(ctx context.Context, req *protoProperty.MsgProperty) (*protoSdk.BaseResponse, error) {
	propertyType := enum.PropertyTypeValue(req.PropertyType)
	amenities := util.ConvertSlice[*model.Amenity](req.Amenities)

	property := &model.Property{
		HostId:       req.HostId,
		PropertyType: &propertyType,
		Status:       &enum.PropertyStatus.InReview,
		OverallRate:  req.OverallRate,

		MaxGuests:    req.MaxGuests,
		MaxPets:      req.MaxPets,
		NumBeds:      req.NumBeds,
		NumBedrooms:  req.NumBedrooms,
		NumBathrooms: req.NumBathrooms,

		IsGuestFavor:  req.IsGuestFavor,
		IsAllowPet:    req.IsAllowPet,
		IsSelfCheckIn: req.IsSelfCheckIn,
		IsInstantBook: req.IsInstantBook,

		Title: req.Title,
		Body:  req.Body,

		NationCode: req.NationCode,
		CityCode:   req.CityCode,
		Lat:        req.Lat,
		Long:       req.Long,

		NightPrice: req.NightPrice,
		ServiceFee: req.ServiceFee,

		Amenities: amenities,
	}

	result := model.PropertyDB.Create(property)

	data := result.Data.([]*model.Property)[0]

	// Sync data to search service
	go bc.SyncProperty(data)

	return util.ConvertToGRPC(result)
}

func (bc *PropertyController) UpdateProperty(ctx context.Context, req *protoProperty.MsgProperty) (*protoSdk.BaseResponse, error) {
	property := &model.Property{
		ID: *req.Id,
	}
	propertyType := enum.PropertyTypeValue(req.PropertyType)
	status := enum.PropertyStatusValue(*req.Status)
	amenities := util.ConvertSlice[model.Amenity](req.Amenities)

	propertyUpdated := &model.Property{
		HostId:       req.HostId,
		PropertyType: &propertyType,
		Status:       &status,
		OverallRate:  req.OverallRate,

		MaxGuests:    req.MaxGuests,
		MaxPets:      req.MaxPets,
		NumBeds:      req.NumBeds,
		NumBedrooms:  req.NumBedrooms,
		NumBathrooms: req.NumBathrooms,

		IsGuestFavor:  req.IsGuestFavor,
		IsAllowPet:    req.IsAllowPet,
		IsSelfCheckIn: req.IsSelfCheckIn,
		IsInstantBook: req.IsInstantBook,

		Title: req.Title,
		Body:  req.Body,

		NationCode: req.NationCode,
		CityCode:   req.CityCode,
		Lat:        req.Lat,
		Long:       req.Long,

		NightPrice: req.NightPrice,
		ServiceFee: req.ServiceFee,
		TaxPercent: req.TaxPercent,
	}

	// BUG HERE: cannot update many2many, fix this shit
	model.PropertyDB.DB.Model(model.PropertyDB.Model).Where(property).Association("Amenities").
		Replace(amenities)

	result := model.PropertyDB.Update(property, propertyUpdated)
	go bc.SyncUpdateProperty(*req.Id)

	return util.ConvertToGRPC(result)
}

func (bc *PropertyController) DeleteProperty(ctx context.Context, req *protoProperty.MsgDeleteProperty) (*protoSdk.BaseResponse, error) {
	propertyId := req.PropertyId
	property := &model.Property{
		ID: propertyId,
	}

	result := model.PropertyDB.Delete(property)
	return util.ConvertToGRPC(result)
}
