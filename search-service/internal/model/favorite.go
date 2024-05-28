package model

import (
	"time"
)

type Favorite struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//foreign key
	PropertyId string `json:"propertyId"  gorm:"column:property_id"`
	UserId     string `json:"userId"  gorm:"column:user_id"`
}
