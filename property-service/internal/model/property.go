package model

import (
	"property-service/internal/model/enum"
	"time"

	"github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type Property struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//has many
	Reviews   []*Review   `json:"reviews,omitempty" gorm:"foreignKey:property_id"`
	Bookings  []*Booking  `json:"bookings,omitempty" gorm:"foreignKey:property_id"`
	Favorites []*Favorite `json:"favorites,omitempty" gorm:"foreignKey:property_id"`

	//many2many
	Amenities []*Amenity `json:"amenities,omitempty" gorm:"many2many:property_amenity;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	//foreign key
	HostId string `json:"hostId,omitempty" gorm:"column:host_id"`

	HostName string `json:"hostName,omitempty" gorm:"column:host_name"`
	HostAvatar  string `json:"hostUrl,omitempty" gorm:"column:host_avatar"`

	PropertyType *enum.PropertyTypeValue   `json:"propertyType,omitempty" gorm:"column:property_type"`
	Status       *enum.PropertyStatusValue `json:"status,omitempty" gorm:"column:status"`

	OverallRate float32 `json:"overallRate,omitempty" gorm:"column:overall_rate"`

	MaxNights    int32 `json:"maxNights,omitempty" gorm:"column:max_nights"`
	MaxGuests    int32 `json:"maxGuests,omitempty" gorm:"column:max_guests"`
	MaxPets      int32 `json:"maxPets,omitempty" gorm:"column:max_pets"`
	NumBeds      int32 `json:"numBeds,omitempty" gorm:"column:num_beds"`
	NumBedrooms  int32 `json:"numBedrooms,omitempty" gorm:"column:num_bedrooms"`
	NumBathrooms int32 `json:"numBathrooms,omitempty" gorm:"column:num_bathrooms"`

	//gorm not return false with bool
	IsGuestFavor  *bool `json:"isGuestFavor,omitempty" gorm:"column:is_guest_favor"`
	IsAllowPet    *bool `json:"isAllowPet,omitempty" gorm:"column:is_allow_pet"`
	IsSelfCheckIn *bool `json:"isSelfCheckIn,omitempty" gorm:"column:is_self_check_in"`
	IsInstantBook *bool `json:"isInstantBook,omitempty" gorm:"column:is_instant_book"`

	Title        string `json:"title,omitempty" gorm:"column:title"`
	Body         string `json:"body,omitempty" gorm:"column:body"`
	Neighborhood string `json:"neighborhood,omitempty" gorm:"column:neighborhood"`

	Address    *string `json:"address,omitempty" gorm:"column:address"`
	CityCode   *string `json:"cityCode,omitempty" gorm:"column:city_code"`
	NationCode *string `json:"nationCode,omitempty" gorm:"column:nation_code"`
	Lat        *string `json:"lat,omitempty" gorm:"column:lat"`
	Long       *string `json:"long,omitempty" gorm:"column:long"`

	NightPrice float64 `json:"nightPrice,omitempty" gorm:"column:night_price"`
	ServiceFee float64 `json:"serviceFee,omitempty" gorm:"column:service_fee"`
	TaxPercent float64 `json:"taxPercent,omitempty" gorm:"column:tax_percent"`

	IntroCover  *string `json:"introCover,omitempty" gorm:"column:intro_cover"`
	IntroImages *string `json:"introImages,omitempty" gorm:"column:intro_images;default:'[]'"`

	BedroomCover  *string `json:"bedroomCover,omitempty" gorm:"column:bedroom_cover"`
	BedroomImages *string `json:"bedroomImages,omitempty" gorm:"column:bedroom_images;default:'[]'"`

	OtherCover  *string `json:"otherCover,omitempty" gorm:"column:other_cover"`
	OtherImages *string `json:"otherImages,omitempty" gorm:"column:other_images;default:'[]'"`
}

func (Property) TableName() string {
	return "property"
}

var PropertyDB = &orm.Instance{
	TableName: "property",
	Model:     &Property{},
}

func InitTableProperty(db *gorm.DB) {
	db.AutoMigrate(&Property{})
	PropertyDB.ApplyDatabase(db)
}
