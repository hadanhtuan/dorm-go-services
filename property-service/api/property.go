package apiProperty

import (
	"context"
	"encoding/json"
	"property-service/internal/model"
	"property-service/internal/model/enum"
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"
	"sync"

	"github.com/hadanhtuan/go-sdk/common"
	"github.com/hadanhtuan/go-sdk/db/orm"
)

func (bc *PropertyAPI) GetProperty(ctx context.Context, req *protoProperty.MsgQueryProperty) (*protoSdk.BaseResponse, error) {
	filter := &model.Property{}

	if req.QueryFields.Id != nil {
		filter.ID = *req.QueryFields.Id
	}

	if req.QueryFields.HostId != "" {
		filter.HostId = req.QueryFields.HostId
	}

	if req.QueryFields.Status != nil && *req.QueryFields.Status != "" {
		status := enum.PropertyStatusValue(*req.QueryFields.Status)
		filter.Status = &status
	}

	if req.QueryFields.PropertyType != "" {
		propertyType := enum.PropertyTypeValue(req.QueryFields.PropertyType)
		filter.PropertyType = &propertyType
	}

	result := model.PropertyDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, &orm.QueryOption{
		Preload: []string{"Amenities", "Bookings"}, //Field name, not table name
		Order:   []string{"created_at desc"},
	})
	result.Message = "Get properties successfully"
	data := result.Data.([]*model.Property)

	for _, item := range data {
		go bc.SyncProperty(item)
	}
	return util.ConvertToGRPC(result)
}

func (bc *PropertyAPI) CountPropertyStatus(ctx context.Context, req *protoProperty.MsgProperty) (*protoSdk.BaseResponse, error) {

	filter := &model.Property{}
	statuses := util.ConvertEnumToSlice(*enum.PropertyStatus)

	if req.Status != nil && *req.Status != "" {
		status := enum.PropertyStatusValue(*req.Status)
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

func (bc *PropertyAPI) CreateProperty(ctx context.Context, req *protoProperty.MsgProperty) (*protoSdk.BaseResponse, error) {
	propertyType := enum.PropertyTypeValue(req.PropertyType)
	amenities := util.ConvertSlice[*model.Amenity](req.Amenities)

	property := &model.Property{
		HostId:       req.HostId,
		HostName:     req.HostName,
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

		Address:    req.Address,
		NationCode: req.NationCode,
		CityCode:   req.CityCode,
		Lat:        req.Lat,
		Long:       req.Long,

		NightPrice: req.NightPrice,
		ServiceFee: req.ServiceFee,

		Amenities: amenities,
	}

	if req.IntroCover != nil && *req.IntroCover != "" {
		property.IntroCover = req.IntroCover
	}

	if len(req.IntroImages) > 0 {
		encode, _ := json.Marshal(req.IntroImages)
		introImages := string(encode)
		property.IntroImages = &introImages
	}

	result := model.PropertyDB.Create(property)

	data := result.Data.([]*model.Property)[0]

	// Sync data to search service
	go bc.SyncProperty(data)

	return util.ConvertToGRPC(result)
}

func (bc *PropertyAPI) UpdateProperty(ctx context.Context, req *protoProperty.MsgProperty) (*protoSdk.BaseResponse, error) {
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

func (bc *PropertyAPI) DeleteProperty(ctx context.Context, req *protoProperty.MsgDeleteProperty) (*protoSdk.BaseResponse, error) {
	propertyId := req.PropertyId
	property := &model.Property{
		ID: propertyId,
	}

	result := model.PropertyDB.Delete(property)
	return util.ConvertToGRPC(result)
}

// func (bc *PropertyAPI) MapMockData(limit int32, offset int32) {
// 	result := model.MockPropertyDB.Query(nil, offset, limit, nil)
// 	data := result.Data.([]*model.MockProperty)

// 	for _, item := range data {
// 		propertyType := enum.PropertyTypeValue(item.PropertyType)
// 		f := false
// 		t := true
// 		lat := fmt.Sprintf("%f", item.Latitude)
// 		long := fmt.Sprintf("%f", item.Longitude)
// 		isInstantBook := true
// 		if item.InstantBookable == "f" {
// 			isInstantBook = false
// 		}

// 		trimString := strings.TrimPrefix(item.Price, "$")
// 		nightPrice, _ := strconv.ParseFloat(trimString, 64)

// 		property := &model.Property{
// 			HostId:        string(item.HostId),
// 			HostFirstName: item.HostName,
// 			HostLastName:  item.HostName,
// 			HostUrl:       item.HostUrl,
// 			PropertyType:  &propertyType,
// 			Status:        &enum.PropertyStatus.InReview,
// 			// OverallRate:   item.OverallRate,

// 			MaxGuests:    int32(item.Accommodates),
// 			MaxPets:      2,
// 			NumBeds:      int32(item.Beds),
// 			NumBedrooms:  int32(item.Bedrooms),
// 			NumBathrooms: int32(item.Bathrooms),

// 			IsGuestFavor:  &f,
// 			IsAllowPet:    &t,
// 			IsSelfCheckIn: &t,
// 			IsInstantBook: &isInstantBook,

// 			Title: item.Name,
// 			Body:  item.Description,

// 			Address: &item.HostLocation,
// 			// NationCode: item.NationCode,
// 			// CityCode:   item.CityCode,
// 			Lat:  &lat,
// 			Long: &long,

// 			NightPrice: nightPrice,
// 			ServiceFee: 15,

// 			// Amenities: amenities,

// 			IntroCover: &item.PictureUrl,
// 		}
// 	}
// }
