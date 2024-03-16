package model

import (
	"time"

	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type Amenity struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	Name        string `json:"name"  gorm:"column:name"`
	Description string `json:"description"  gorm:"column:description"`
	Icon        string `json:"icon"  gorm:"column:icon"`
}

// Overwrite table name
func (Amenity) TableName() string {
	return "amenity"
}

var AmenityDB = &orm.Instance{
	TableName: "amenity",
	Model:     &Amenity{},
}

func InitTableAmenity(db *gorm.DB) {
	db.AutoMigrate(&Amenity{})
	AmenityDB.ApplyDatabase(db)
}
