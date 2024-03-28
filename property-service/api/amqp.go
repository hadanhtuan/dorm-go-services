package apiProperty

import (
	"encoding/json"
	"fmt"
	"property-service/internal/model"
	"property-service/internal/util"

	"github.com/hadanhtuan/go-sdk/amqp"
	"github.com/hadanhtuan/go-sdk/db/orm"
)

// Sync data to Search Service
func (bc *PropertyController) SyncProperty(data *model.Property) {
	encodeData, _ := json.Marshal(data)
	instant := amqp.GetConnection()
	err := instant.PublishMessage(util.EXCHANGE, util.PropertyCreated, encodeData)

	if err != nil {
		fmt.Println(err.Error())
	}
}

// Sync data to Search Service
func (bc *PropertyController) SyncUpdateProperty(id string) {
	filter := &model.Property{
		ID: id,
	}
	result := model.PropertyDB.QueryOne(filter, &orm.QueryOption{
		Preload: []string{"Amenities", "Bookings"},
	})

	data := result.Data.([]*model.Property)[0]
	encodeData, _ := json.Marshal(data)

	instant := amqp.GetConnection()
	err := instant.PublishMessage(util.EXCHANGE, util.PropertyUpdated, encodeData)

	if err != nil {
		fmt.Println(err.Error())
	}
}
