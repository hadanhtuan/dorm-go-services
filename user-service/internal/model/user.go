package model

import (
	// "time"

	"github.com/google/uuid"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	// CreatedAt time.Time  `json:"createdAt,omitempty"`
	// UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	// DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Username  string     `json:"username,omitempty"`
	// FullName  string     `json:"fullName,omitempty"`
	// Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	// IsActive  bool       `json:"isActive,omitempty"`
}

var UserDB = &orm.Instance{
	TableName: "user",
	Model:     &User{},
}

func InitTableUser(db *gorm.DB) {
	UserDB.ApplyDatabase(db)
}
