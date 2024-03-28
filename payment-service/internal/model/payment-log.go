package model

import (
	"time"

	"github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type PaymentLog struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	StripeId   *string `json:"stripeId"  gorm:"column:stripe_id"`
	UserId     *string `json:"userId"  gorm:"column:user_id"`
	BookingId  *string `json:"bookingId"  gorm:"column:booking_id"`
	PropertyId *string `json:"propertyId"  gorm:"column:property_id"`

	Amount        int64  `json:"amount"  gorm:"column:amount"`
	Currency      string `json:"currency"  gorm:"column:currency"`
	Event         string `json:"event"  gorm:"column:event"`
	PaymentMethod string `json:"paymentMethod"  gorm:"column:payment_method"`
	Status        string `json:"status"  gorm:"column:status"`
	ReceiptEmail  string `json:"receiptEmail"  gorm:"column:receipt_email"`
	CanceledAt    *int64 `json:"canceledAt"  gorm:"column:canceled_at"`
	Customer      string `json:"customer"  gorm:"column:customer"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

var PaymentLogDB = &orm.Instance{
	TableName: "payment_log",
	Model:     &PaymentLog{},
}

func InitTablePaymentLog(db *gorm.DB) {
	db.AutoMigrate(&PaymentLog{})
	PaymentLogDB.ApplyDatabase(db)
}
