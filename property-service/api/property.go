package apiProperty

import (
	"context"
	"fmt"
	"property-service/internal/model"
	"property-service/internal/model/enum"
	"property-service/internal/util"
	protoProperty "property-service/proto/property"
	protoSdk "property-service/proto/sdk"

	"github.com/hadanhtuan/go-sdk/db/orm"
)

func (bc *PropertyController) GetProperty(ctx context.Context, req *protoProperty.MsgQueryProperty) (*protoSdk.BaseResponse, error) {
	filter := &model.Property{}

	if req.QueryFields.Id != nil {
		filter.ID = *req.QueryFields.Id
	}

	result := model.PropertyDB.Query(filter, req.Paginate.Offset, req.Paginate.Limit, &orm.QueryOption{
		Preload: []string{"Reviews", "Amenities"}, //Field name, not table name
		Order:   []string{"created_at desc"},
	})

	result.Message = "Get properties successfully"
	return util.ConvertToGRPC(result)
}

func (bc *PropertyController) CreateProperty(ctx context.Context, req *protoProperty.MsgProperty) (*protoSdk.BaseResponse, error) {
	propertyType := enum.PropertyTypeValue(req.PropertyType)

	fmt.Println(req.IsAllowPet)
	amenities := util.ConvertSlice[*model.Amenity](req.Amenities)

	property := &model.Property{
		HostId:       req.HostId,
		PropertyType: &propertyType,
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

	amenities := util.ConvertSlice[model.Amenity](req.Amenities)

	propertyUpdated := &model.Property{
		HostId:       req.HostId,
		PropertyType: &propertyType,
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
