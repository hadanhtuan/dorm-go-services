package model

import (
	"booking-service/internal/model/enum"
	"time"

	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type Property struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	ImageUrl      string  `json:"imageUrl,omitempty" gorm:"column:image_url"`
	OverallRating float64 `json:"overallRate,omitempty" gorm:"column:overall_rate"`

	WardId string `json:"wardId" gorm:"column:ward_id"`
	Lat    string `json:"lat" gorm:"column:lat"`
	Long   string `json:"long" gorm:"column:long"`
	HostId string `json:"hostId" gorm:"column:host_id"`

	Amenities string `json:"amenities" gorm:"column:amenities"`

	PropertyType *enum.PropertyTypeValue `json:"propertyType,omitempty" gorm:"column:property_type"`
	NumGuests    int32                   `json:"numGuests,omitempty" gorm:"column:num_guests"`
	NumBeds      int32                   `json:"numBeds,omitempty" gorm:"column:num_beds"`
	NumBedrooms  int32                   `json:"numBedrooms,omitempty" gorm:"column:num_bedrooms"`
	NumBaths     int32                   `json:"numBathrooms,omitempty" gorm:"column:num_bathrooms"`
	Price        float64                 `json:"price,omitempty" gorm:"column:price"`
	IsGuestFavor bool                    `json:"isGuestFavor,omitempty" gorm:"column:is_guest_favor"`
	Body         string                  `json:"body,omitempty" gorm:"column:body"`
	Title        string                  `json:"title,omitempty" gorm:"column:title"`
	// ImageUrl     []*string               `json:"imageUrl,omitempty" gorm:"column:image_url"`
}

var PropertyDB = &orm.Instance{
	TableName: "properties",
	Model:     &Property{},
}

func InitTableProperty(db *gorm.DB) {
	db.Table(PropertyDB.TableName).AutoMigrate(&Property{})
	PropertyDB.ApplyDatabase(db)
}
