package model

import (
	"property-service/internal/model/enum"
	"time"

	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type Booking struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//foreign key
	PropertyId string `json:"propertyId"  gorm:"column:property_id"`
	UserId     string `json:"userId"  gorm:"column:user_id"`

	CheckInDate  time.Time `json:"checkInDate,omitempty" gorm:"column:checkIn_date"`
	CheckoutDate time.Time `json:"checkoutDate,omitempty" gorm:"column:checkout_date"`

	DiscountFee float32 `json:"discountFee,omitempty" gorm:"column:discount_fee"`
	TotalPrice  float32 `json:"totalPrice,omitempty" gorm:"column:total_price"`
	
	GuestNumber  int32                    `json:"guestNumber,omitempty" gorm:"column:guest_number"`
	ChildNumber  int32                    `json:"childNumber,omitempty" gorm:"column:child_number"`
	BabyNumber   int32                    `json:"babyNumber,omitempty" gorm:"column:baby_number"`
	PetNumber    int32                    `json:"petNumber,omitempty" gorm:"column:pet_number"`
	Status       *enum.BookingStatusValue `json:"status,omitempty" gorm:"column:status"`
	HostLanguage string                   `json:"hostLanguage,omitempty" gorm:"column:host_language"`
}

func (Booking) TableName() string {
	return "booking"
}

var BookingDB = &orm.Instance{
	TableName: "booking",
	Model:     &Booking{},
}

func InitTableBooking(db *gorm.DB) {
	db.AutoMigrate(&Booking{})
	BookingDB.ApplyDatabase(db)
}
