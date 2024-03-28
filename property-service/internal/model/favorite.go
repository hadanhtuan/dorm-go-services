package model

import (
	"time"

	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
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

// Overwrite table name
func (Favorite) TableName() string {
	return "favorite"
}

var FavoriteDB = &orm.Instance{
	TableName: "favorite",
	Model:     &Favorite{},
}

func InitTableFavorite(db *gorm.DB) {
	db.AutoMigrate(&Favorite{})
	FavoriteDB.ApplyDatabase(db)
}
