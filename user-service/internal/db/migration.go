package database

import (
	"fmt"
	"user-service/internal/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	fmt.Printf("Migrate: Start")
	err := db.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Migrate: Success")
}
