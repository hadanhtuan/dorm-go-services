package apiProperty

import (
	"encoding/json"
	"fmt"
	"property-service/internal/model"
	"property-service/internal/model/enum"
	"property-service/internal/util"
	"time"

	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/hadanhtuan/go-sdk/db/orm"
)

func (pc *PropertyAPI) bindingMap() map[string]amqp.CallbackFunc {
	return map[string]amqp.CallbackFunc{
		util.PaymentSuccess: pc.EventPaymentSuccess,
	}
}

func (pc *PropertyAPI) InitRoutingAMQP() {
	instance := amqp.GetConnection()
	bindingMap := pc.bindingMap()

	for routingKey, execFunc := range bindingMap {
		instance.BindingQueue(util.PROPERTY_EXCHANGE, util.PROPERTY_QUEUE, routingKey, execFunc)
	}

	instance.StartConsume(util.PROPERTY_QUEUE)
}

// Receive data from payment Service
func (pc *PropertyAPI) EventPaymentSuccess(payload []byte) {
	var data model.Booking
	json.Unmarshal(payload, &data)

	filter := data

	data.PaymentDate = time.Now().Unix()
	data.Status = &enum.BookingStatus.Success
	_ = model.BookingDB.Update(filter, data)
}

// Sync data to Search Service
func (bc *PropertyAPI) SyncProperty(data *model.Property) {
	encodeData, _ := json.Marshal(data)
	instant := amqp.GetConnection()
	err := instant.PublishMessage(util.SEARCH_EXCHANGE, util.PropertyCreated, encodeData)

	if err != nil {
		fmt.Println(err.Error())
	}
}

// Sync data to Search Service
func (bc *PropertyAPI) SyncUpdateProperty(id string) {
	filter := &model.Property{
		ID: id,
	}
	result := model.PropertyDB.QueryOne(filter, &orm.QueryOption{
		Preload: []string{"Amenities", "Bookings"},
	})

	if result.Status == common.APIStatus.NotFound {
		return
	}

	data := result.Data.([]*model.Property)[0]


	fmt.Println(*data.Status)
	encodeData, _ := json.Marshal(&data)

	instant := amqp.GetConnection()
	err := instant.PublishMessage(util.SEARCH_EXCHANGE, util.PropertyUpdated, encodeData)

	if err != nil {
		fmt.Println(err.Error())
	}
}
