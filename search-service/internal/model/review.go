package model

import (
	"time"
)

type Review struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//foreign key
	PropertyId string  `json:"propertyId" gorm:"column:property_id"`
	ParentId   *string `json:"parentId"  gorm:"column:parent_id"`
	UserId     string  `json:"userId"  gorm:"column:user_id"`

	Rating   float32 `json:"overallRating"  gorm:"column:overall_rating"`
	Comment  string  `json:"comment"  gorm:"column:comment"`
	ImageUrl string  `json:"imageUrl,omitempty" gorm:"column:image_url"`
}
