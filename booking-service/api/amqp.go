package apiBooking

import (
	"booking-service/internal/model"
	"booking-service/internal/util"
	"encoding/json"
	"fmt"

	"github.com/hadanhtuan/go-sdk/amqp"
)

// Publish Message
func (bc *BookingController) SyncProperty(data *model.Property) {
	encodeData, _ := json.Marshal(data)
	instant := amqp.GetConnection()
	err := instant.PublishMessage(util.EXCHANGE, util.PropertyCreated, encodeData)

	if err != nil {
		fmt.Println(err.Error())
	}
}
