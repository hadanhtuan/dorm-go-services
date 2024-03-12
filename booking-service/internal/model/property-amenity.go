package model

import (
	"time"

	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type PropertyAmenity struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	PropertyId string `json:"PropertyId"  gorm:"column:property_id"`
	AmenityId  string `json:"AmenityId"  gorm:"column:amenity_id"`
	Order      int32  `json:"order"  gorm:"column:order"`
}

var PropertyAmenityDB = &orm.Instance{
	TableName: "property_amenity",
	Model:     &PropertyAmenity{},
}

func InitTablePropertyAmenity(db *gorm.DB) {
	db.Table(PropertyAmenityDB.TableName).AutoMigrate(&PropertyAmenity{})
	PropertyAmenityDB.ApplyDatabase(db)
}
