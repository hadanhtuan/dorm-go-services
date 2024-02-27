package model

import (
	"time"
	"user-service/internal/model/enum"

	"github.com/google/uuid"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	Role *enum.UserRoleValue `json:"role,omitempty" gorm:"column:role"`

	Username  string `json:"username,omitempty" gorm:"column:username"`
	FirstName string `json:"firstName,omitempty" gorm:"column:first_name"`
	LastName  string `json:"lastName,omitempty" gorm:"column:last_name"`
	Email     string `json:"email,omitempty" gorm:"column:email"`
	Phone     string `json:"phone,omitempty" gorm:"column:phone"`
	Password  string `json:"password,omitempty" gorm:"column:password"`
	IsActive  bool   `json:"isActive,omitempty" gorm:"column:is_active"`
}

var UserDB = &orm.Instance{
	TableName: "member",
	Model:     &User{},
}

func InitTableUser(db *gorm.DB) {
	db.Table(UserDB.TableName).AutoMigrate(&User{})
	UserDB.ApplyDatabase(db)
}
