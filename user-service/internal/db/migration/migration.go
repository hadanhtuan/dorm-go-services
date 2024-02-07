package database

import (
	"fmt"
	"user-service/internal/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	fmt.Printf("Migrate: Start")
	db.AutoMigrate(
		&model.User{},
	)
	fmt.Printf("Migrate: Success")
}
