package model

import (
	"time"
)

type Amenity struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	Name        string `json:"name"  gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Icon        string `json:"icon"  gorm:"column:icon"`
}
