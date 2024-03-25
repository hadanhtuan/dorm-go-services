package model

import (
	"search-service/internal/model/enum"
	"time"
)

type Booking struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//foreign key
	PropertyId string `json:"propertyId"  gorm:"column:property_id"`
	UserId     string `json:"userId"  gorm:"column:user_id"`

	CheckInDate   time.Time                `json:"checkInDate,omitempty" gorm:"column:checkIn_date"`
	CheckoutDate  time.Time                `json:"checkoutDate,omitempty" gorm:"column:checkout_date"`
	TotalPrice    float64                  `json:"totalPrice,omitempty" gorm:"column:total_price"`
	ServiceFee    float64                  `json:"serviceFee,omitempty" gorm:"column:service_fee"`
	GuestNumber   int32                    `json:"guestNumber,omitempty" gorm:"column:guest_number"`
	ChildNumber   int32                    `json:"childNumber,omitempty" gorm:"column:child_number"`
	BabyNumber    int32                    `json:"babyNumber,omitempty" gorm:"column:baby_number"`
	PetNumber     int32                    `json:"petNumber,omitempty" gorm:"column:pet_number"`
	IsInstantBook bool                     `json:"isInstantBook,omitempty" gorm:"column:is_instant_book"`
	IsSelfCheckIn bool                     `json:"isSelfCheckIn,omitempty" gorm:"column:is_self_checkIn"`
	IsAllowPet    bool                     `json:"isAllowPet,omitempty" gorm:"column:is_allow_pet"`
	Status        *enum.BookingStatusValue `json:"status,omitempty" gorm:"column:status"`
	HostLanguage  string                   `json:"hostLanguage,omitempty" gorm:"column:host_language"`
}
