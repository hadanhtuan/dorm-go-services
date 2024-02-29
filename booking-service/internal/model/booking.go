package model

import (
	"booking-service/internal/model/enum"
	"time"

	"github.com/google/uuid"
	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type Booking struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	// Role *enum.UserRoleValue `json:"role,omitempty" gorm:"column:role"`
	UserId        uuid.UUID                `json:"userId"  gorm:"type: uuid;column:user_id"`
	PropertyId    uuid.UUID                `json:"propertyId"  gorm:"type: uuid;column:property_id"`
	CheckinDate   time.Time                `json:"checkinDate,omitempty" gorm:"column:checkin_date"`
	CheckoutDate  time.Time                `json:"checkoutDate,omitempty" gorm:"column:checkout_date"`
	TotolPrice    float64                  `json:"totolPrice,omitempty" gorm:"column:totol_price"`
	ServiceFee    float64                  `json:"serviceFee,omitempty" gorm:"column:service_fee"`
	GuestNumber   int                      `json:"guestNumber,omitempty" gorm:"column:guest_number"`
	ChildNumber   int                      `json:"childNumber,omitempty" gorm:"column:child_number"`
	BabyNumber    int                      `json:"babyNumber,omitempty" gorm:"column:baby_number"`
	PetNumber     int                      `json:"petNumber,omitempty" gorm:"column:pet_number"`
	IsInstantBook bool                     `json:"isInstantBook,omitempty" gorm:"column:is_instant_book"`
	IsSelfCheckin bool                     `json:"isSelfCheckin,omitempty" gorm:"column:is_self_checkin"`
	IsAllowPet    bool                     `json:"isAllowPet,omitempty" gorm:"column:is_allow_pet"`
	Status        *enum.BookingStatusValue `json:"status,omitempty" gorm:"column:status"`
	HostLanguage  string                   `json:"hostLanguage,omitempty" gorm:"column:host_language"`
}

var BookingDB = &orm.Instance{
	TableName: "booking",
	Model:     &Booking{},
}

func InitTableUser(db *gorm.DB) {
	db.Table(BookingDB.TableName).AutoMigrate(&Booking{})
	BookingDB.ApplyDatabase(db)
}
