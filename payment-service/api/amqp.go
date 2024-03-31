package apiPayment

import (
	"encoding/json"
	"fmt"
	"payment-service/internal/model"
	"payment-service/internal/util"

	"github.com/hadanhtuan/go-sdk/amqp"
)

// Sync data to Search Service
func (pc *PaymentController) BookingSuccess(bookingId string) {
	data := &model.Booking{
		ID: bookingId,
	}
	encodeData, _ := json.Marshal(data)
	instant := amqp.GetConnection()
	err := instant.PublishMessage(util.PROPERTY_EXCHANGE, util.PAYMENT_EXCHANGE, encodeData)

	if err != nil {
		fmt.Println(err.Error())
	}
}
