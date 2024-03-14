package model

import (
	"time"

	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
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

func (Review) TableName() string {
	return "review"
}

var ReviewDB = &orm.Instance{
	TableName: "review",
	Model:     &Review{},
}

func InitTableReview(db *gorm.DB) {
	db.AutoMigrate(&Review{})
	ReviewDB.ApplyDatabase(db)
}
