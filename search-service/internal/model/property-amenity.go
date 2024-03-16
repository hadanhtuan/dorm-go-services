package model

import (
	"time"
)

type PropertyAmenity struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	PropertyID string `gorm:"primaryKey,column:property_id"`
	AmenityID  string `gorm:"primaryKey,column:amenity_id"`

	Order int32 `json:"order"  gorm:"column:order"`
}
