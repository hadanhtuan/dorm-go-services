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
	Reviews           []*Review          `json:"reviews,omitempty" gorm:"foreignKey:property_id"`
	PropertyAmenities []*PropertyAmenity `json:"propertyAmenities,omitempty" gorm:"foreignKey:property_id"`

	PropertyType *enum.PropertyTypeValue `json:"propertyType,omitempty" gorm:"column:property_type"`
	ImageUrl     string                  `json:"imageUrl,omitempty" gorm:"column:image_url"`
	OverallRate  float32                 `json:"overallRate,omitempty" gorm:"column:overall_rate"`
	WardId       string                  `json:"wardId" gorm:"column:ward_id"`
	Lat          string                  `json:"lat" gorm:"column:lat"`
	Long         string                  `json:"long" gorm:"column:long"`
	HostId       string                  `json:"hostId" gorm:"column:host_id"`
	NumGuests    int32                   `json:"numGuests,omitempty" gorm:"column:num_guests"`
	NumBeds      int32                   `json:"numBeds,omitempty" gorm:"column:num_beds"`
	NumBedrooms  int32                   `json:"numBedrooms,omitempty" gorm:"column:num_bedrooms"`
	NumBaths     int32                   `json:"numBathrooms,omitempty" gorm:"column:num_bathrooms"`
	Price        float32                 `json:"price,omitempty" gorm:"column:price"`
	IsGuestFavor bool                    `json:"isGuestFavor,omitempty" gorm:"column:is_guest_favor"`
	Body         string                  `json:"body,omitempty" gorm:"column:body"`
	Title        string                  `json:"title,omitempty" gorm:"column:title"`
}
