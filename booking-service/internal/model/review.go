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

	UserId     string  `json:"userId"  gorm:"column:user_id"`
	PropertyId string  `json:"propertyId"  gorm:"column:property_id"`
	ParentId   string  `json:"parentId"  gorm:"column:parent_id"`
	Rating     float32 `json:"overallRating"  gorm:"column:overall_rating"`
	Comment    string  `json:"comment"  gorm:"column:comment"`
	ImageUrl   string  `json:"imageUrl,omitempty" gorm:"column:image_url"`
}

var ReviewDB = &orm.Instance{
	TableName: "reviews",
	Model:     &Review{},
}

func InitTableReview(db *gorm.DB) {
	db.Table(ReviewDB.TableName).AutoMigrate(&Review{})
	ReviewDB.ApplyDatabase(db)
}
