package apiBooking

import (
	"booking-service/internal/model"
	"booking-service/internal/util"
	"encoding/json"

	"github.com/hadanhtuan/go-sdk/amqp"
)

func (bc *BookingController) SyncProperty(data *model.Property) {
	encodeData, _ := json.Marshal(data)
	client := amqp.GetConnection()
	client.PublishMessage(util.EXCHANGE, util.PropertyCreated, encodeData)
}
