package model

import (
	"search-service/internal/model/enum"
	"time"
)

type Property struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//has many
	Reviews  []*Review  `json:"reviews,omitempty" gorm:"foreignKey:property_id"`
	Bookings []*Booking `json:"bookings,omitempty" gorm:"foreignKey:property_id"`

	//many2many
	Amenities []*Amenity `json:"amenities,omitempty" gorm:"many2many:property_amenity;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	//foreign key
	HostId string `json:"hostId,omitempty" gorm:"column:host_id"`

	PropertyType *enum.PropertyTypeValue `json:"propertyType,omitempty" gorm:"column:property_type"`
	OverallRate  float32                 `json:"overallRate,omitempty" gorm:"column:overall_rate"`

	MaxGuests    int32 `json:"maxGuests,omitempty" gorm:"column:max_guests"`
	MaxPets      int32 `json:"maxPets,omitempty" gorm:"column:max_Pets"`
	NumBeds      int32 `json:"numBeds,omitempty" gorm:"column:num_beds"`
	NumBedrooms  int32 `json:"numBedrooms,omitempty" gorm:"column:num_bedrooms"`
	NumBathrooms int32 `json:"numBathrooms,omitempty" gorm:"column:num_bathrooms"`

	//gorm not return false with bool
	IsGuestFavor  *bool `json:"isGuestFavor,omitempty" gorm:"column:is_guest_favor"`
	IsAllowPet    *bool `json:"isAllowPet,omitempty" gorm:"column:is_allow_pet"`
	IsSelfCheckIn *bool `json:"isSelfCheckIn,omitempty" gorm:"column:is_self_check_in"`
	IsInstantBook *bool `json:"isInstantBook,omitempty" gorm:"column:is_instant_book"`

	Title string `json:"title,omitempty" gorm:"column:title"`
	Body  string `json:"body,omitempty" gorm:"column:body"`

	CityCode   *string `json:"cityCode,omitempty" gorm:"column:city_code"`
	NationCode *string `json:"nationCode,omitempty" gorm:"column:nation_code"`
	Lat        *string `json:"lat,omitempty" gorm:"column:lat"`
	Long       *string `json:"long,omitempty" gorm:"column:long"`

	NightPrice float32 `json:"nightPrice,omitempty" gorm:"column:night_price"`
	ServiceFee float32 `json:"serviceFee,omitempty" gorm:"column:service_fee"`
}
