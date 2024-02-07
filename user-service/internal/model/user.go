package model

import (
	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CoreModel
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsActive bool   `json:"isActive"`
}

var UserDB = &orm.Instance{
	TableName: "user",
	Model:     &User{},
}

func InitTableUser(db *gorm.DB) {
	UserDB.ApplyDatabase(db)
}
