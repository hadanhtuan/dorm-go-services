package model

import (
	"property-service/internal/model/enum"
	"time"
)

type User struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	Role *enum.UserRoleValue `json:"role,omitempty" gorm:"column:role"`

	Username  string `json:"username,omitempty" gorm:"column:username"`
	// FirstName string `json:"firstName,omitempty" gorm:"column:first_name"`
	// LastName  string `json:"lastName,omitempty" gorm:"column:last_name"`
	Email     string `json:"email,omitempty" gorm:"column:email"`
	Phone     string `json:"phone,omitempty" gorm:"column:phone"`
	Password  string `json:"password,omitempty" gorm:"column:password"`
	Avatar    string `json:"avatar,omitempty" gorm:"column:avatar"`
	Address   string `json:"address,omitempty" gorm:"column:address"`
	Gender    string `json:"gender,omitempty" gorm:"column:gender"`
	DoB       int64  `json:"dob,omitempty" gorm:"column:dob"`
	IsActive  *bool  `json:"isActive,omitempty" gorm:"column:is_active"`
}