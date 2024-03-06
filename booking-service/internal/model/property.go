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

	// Role *enum.UserRoleValue `json:"role,omitempty" gorm:"column:role"`
	WardId       string                  `json:"wardId"  gorm:"column:ward_id"`
	HostId       string                  `json:"hostId"  gorm:"column:host_id"`
	NumGuests    int                     `json:"numGuests,omitempty" gorm:"column:num_guests"`
	NumBeds      int                     `json:"numBeds,omitempty" gorm:"column:num_beds"`
	NumBedrooms  int                     `json:"numBedrooms,omitempty" gorm:"column:num_bedrooms"`
	NumBaths     int                     `json:"numBathrooms,omitempty" gorm:"column:num_bathrooms"`
	IsGuestFavor bool                    `json:"isGuestFavor,omitempty" gorm:"column:is_guest_favor"`
	PropertyType *enum.PropertyTypeValue `json:"propertyType,omitempty" gorm:"column:property_type"`
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
