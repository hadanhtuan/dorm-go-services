package model

import (
	"booking-service/internal/model/enum"
	"time"

	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type Property struct {
	ID        string     `json:"id" gorm:"primarykey;default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//has many
	Reviews  []*Review  `json:"reviews,omitempty" gorm:"foreignKey:property_id"`
	Bookings []*Booking `json:"bookings,omitempty" gorm:"foreignKey:property_id"`

	//many2many
	Amenities []*Amenity `json:"amenities,omitempty" gorm:"many2many:property_amenity;"`

	//foreign key
	HostId string `json:"hostId" gorm:"column:host_id"`

	PropertyType *enum.PropertyTypeValue `json:"propertyType,omitempty" gorm:"column:property_type"`
	OverallRate  float32                 `json:"overallRate,omitempty" gorm:"column:overall_rate"`

	WardCode   *string `json:"wardCode,omitempty" gorm:"column:ward_code"`
	NationCode *string `json:"nationCode,omitempty" gorm:"column:nation_code"`
	Lat        *string `json:"lat,omitempty" gorm:"column:lat"`
	Long       *string `json:"long,omitempty" gorm:"column:long"`

	NumGuests     int32   `json:"numGuests,omitempty" gorm:"column:num_guests"`
	NumBeds       int32   `json:"numBeds,omitempty" gorm:"column:num_beds"`
	NumBedrooms   int32   `json:"numBedrooms,omitempty" gorm:"column:num_bedrooms"`
	NumBathRooms  int32   `json:"numBathrooms,omitempty" gorm:"column:num_bathrooms"`
	Price         float32 `json:"price,omitempty" gorm:"column:price"`
	IsGuestFavor  bool    `json:"isGuestFavor,omitempty" gorm:"column:is_guest_favor"`
	IsAllowPet    bool    `json:"isAllowPet,omitempty" gorm:"column:is_allow_pet"`
	IsSelfCheckIn bool    `json:"isSelfCheckIn,omitempty" gorm:"column:is_self_check_in"`
	IsInstantBook bool    `json:"isInstantBook,omitempty" gorm:"column:is_instant_book"`
	Title         string  `json:"title,omitempty" gorm:"column:title"`
	Body          string  `json:"body,omitempty" gorm:"column:body"`
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
